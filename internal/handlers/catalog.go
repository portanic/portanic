package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/catalog"
	"github.com/portanic/portanic/internal/db"
)

type CatalogHandler struct {
	DB *db.Queries
}

func (h CatalogHandler) HandleShowCataLog(c echo.Context) error {
	ctx := context.Background()
	catalog, err := h.DB.GetCatalog(ctx)

	template, err := h.DB.GetTemplateByID(ctx, catalog.Template)
	if err != nil {
		panic("catalog shiat")
	}

	fields, err := ParseFields(template.Fields)
	if err != nil {
		panic("Failed to parse fields")
	}

	return render(c, Catalog.Show(fields))
}

// Parse dynamic JSON data into a map
func ParseFields(jsonData []byte) (map[string]string, error) {
	var fieldMap map[string]string
	err := json.Unmarshal(jsonData, &fieldMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling fields data: %w", err)
	}
	return fieldMap, nil
}

func (h CatalogHandler) HandleNewCatalog(c echo.Context) error {
	allTemplates, err := h.DB.GetAllTemplates(context.Background())
	if err != nil {
		panic("failed to retrieve")
	}
	templateList := []Catalog.Template{}

	for _, template := range allTemplates {
		tempTemplate := Catalog.Template{
			Id:   template.ID.String(),
			Name: template.Name,
		}
		templateList = append(templateList, tempTemplate)
	}
	return render(c, Catalog.New(templateList))
}

func (h CatalogHandler) HandleCreateCatalog(c echo.Context) error {
	ctx := context.Background()
	if err := c.Request().ParseForm(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse form data")
	}

	fmt.Println(c)
	templateId := c.FormValue("template")
	fmt.Println("Template Name:", templateId)

	id := uuid.New()
	templateUUID, err := uuid.Parse(templateId)
	if err != nil {
		panic("failed to parse uuid")
	}

	args := db.CreateCatalogParams{
		ID:       id,
		Template: templateUUID,
	}

	h.DB.CreateCatalog(ctx, args)

	return c.String(http.StatusOK, "Catalog created successfully")
}
