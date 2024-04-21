package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/home"
)

type HomeHandler struct {}

func (h HomeHandler) HandleHome(c echo.Context) error {
	return render(c, Home.Show())
}