package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/entry"
	"github.com/portanic/portanic/internal/db"
	"github.com/portanic/portanic/internal/types"
)

type EntryHandler struct {
	DB *db.Queries
}

func (h EntryHandler) HandleNewEntry(c echo.Context) error {
	ctx := context.Background()
	catalog, err := h.DB.GetCatalog(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid catalog ID"})
	}

	var entryData map[string]interface{}
	if err := c.Bind(&entryData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data format"})
	}

	// Fetch template ID associated with the catalog ID
	template, err := h.DB.GetTemplateByID(ctx, catalog.Template)
	if err != nil {
		panic("shiat")
	}

	var fieldsMap map[string]string
	err = json.Unmarshal(template.Fields, &fieldsMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data format"})
	}
	if !types.ValidateData(entryData, fieldsMap) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data does not match template"})
	}

	return render(c, Entry.New(fieldsMap))
}

func (h EntryHandler) HandleCreateEntry(c echo.Context) error {
	ctx := context.Background()
	if err := c.Request().ParseForm(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse form data")
	}

	formValues, err := c.FormParams()
	if err != nil {
		return err
	}

	catalog, err := h.DB.GetCatalog(ctx)
	if err != nil {
		panic("Failed to get catalog")
	}

	id := uuid.New()
	catalogUUID, err := uuid.Parse(catalog.ID.String())
	if err != nil {
		panic("failed to parse uuid")
	}

	data := make(map[string]string)
	for key, values := range formValues {
		if len(values) > 0 {
			data[key] = values[0]
		}
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic("failed to marshal json")
	}

	args := db.CreateEntryParams{
		ID:        id,
		CatalogID: catalogUUID,
		Data:      jsonData,
	}

	h.DB.CreateEntry(ctx, args)
	return c.String(http.StatusOK, "Entry created successfully")
}
