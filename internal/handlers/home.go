package handlers

import "github.com/labstack/echo/v4"

type HomeHandler struct {}

func (h HomeHandler) HandleHome(c echo.Context) error {
    return c.File("static/index.html")
}