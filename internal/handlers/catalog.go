package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/catalog"
)

type CatalogHandler struct {}

func (h CatalogHandler) HandleShowCataLog(c echo.Context) error {
	return render(c, Catalog.Show())
}