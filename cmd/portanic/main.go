package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/handlers"
)

func main() {
    app := echo.New()

    homeHandler := handlers.HomeHandler{}
	app.GET("/", homeHandler.HandleHome)

	app.Static("/css", "static/css")
	app.Static("/js", "static/js")
	app.Static("/vendors", "static/vendors")

    fmt.Println("Server is running on http://localhost:8080")
    app.Logger.Fatal(app.Start(":8080"))
}