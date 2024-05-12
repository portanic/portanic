package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/templates"
	"github.com/portanic/portanic/internal/db"
)

type TemplatesHandler struct {
	DB *db.Queries
}

func (h TemplatesHandler) HandleShowTemplates(c echo.Context) error {
	allTemplates, err := h.DB.GetAllTemplates(context.Background())
	if err != nil {
		panic("failed to retrieve")
	}
	templateList := []Templates.Template{}

	for _, template := range allTemplates {
		tempTemplate := Templates.Template{
			Name:      template.Name,
			CreatedAt: template.CreatedAt.Time.String(),
		}
		templateList = append(templateList, tempTemplate)
	}
	return render(c, Templates.Show(templateList))
}

func (h TemplatesHandler) HandleNewTemplate(c echo.Context) error {
	return render(c, Templates.New())
}

func (h TemplatesHandler) HandleCreateTemplate(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse form data")
	}

	templateName := c.FormValue("name")
	fmt.Println("Template Name:", templateName)

	fieldMap := make(map[string]string)

	for key, values := range c.Request().PostForm {
		if strings.HasPrefix(key, "fieldName") {
			index := strings.TrimPrefix(key, "fieldName")
			fieldTypeKey := "fieldType" + index

			fieldName := values[0] // We assume there is at least one value
			fieldType := c.FormValue(fieldTypeKey)
			fmt.Printf("Custom Field Name: %s, Type: %s\n", fieldName, fieldType)

			fieldMap[fieldName] = fieldType
		}
	}

	fieldJSON, err := json.Marshal(fieldMap)
	if err != nil {
		panic("fukc")
	}

	id := uuid.New()

	args := db.CreateTemplateParams{
		ID:        id,
		Name:      templateName,
		Fields:    fieldJSON,
		CreatedBy: "Marius",
	}

	h.DB.CreateTemplate(context.Background(), args)

	return c.String(http.StatusOK, "Template created successfully")
}

func (h TemplatesHandler) HandleGetTemplateFields(c echo.Context) error {
	fmt.Println("Handling get templates")
	templateID := c.QueryParam("tid")
	templateUUID, err := uuid.Parse(templateID)
	if err != nil {
		html := "<div>"
		html += "</div>"
		return c.HTML(http.StatusNotFound, html)
	}

	// Fetch fields based on templateID
	template, err := h.DB.GetTemplateByID(context.Background(), templateUUID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching template fields")
	}

	var typeMap map[string]string
	err = json.Unmarshal(template.Fields, &typeMap)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	// Generate HTML for fields
	html := "<form id='dynamicForm' hx-post='/catalog' hx-target='#formResponse'>"
	for key, value := range typeMap {
		html += fmt.Sprintf("<div><label>%s:</label><input type='text' name='%s' value=''></div>", key, value)
	}
	html += "<button type='submit'>Submit</button>"
	html += "</form>"
	return c.HTML(http.StatusOK, html)
}
