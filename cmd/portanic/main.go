package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/db"
	"github.com/portanic/portanic/internal/handlers"
)

func main() {
    app := echo.New()

	config := db.Config{
		Host:     "db",
		Port:     5432,
		User:     "portanic",
		Password: "portanic",
		DBName:   "portanic",
	}

	database, err := db.New(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

    homeHandler := handlers.HomeHandler{}
	app.GET("/", homeHandler.HandleHome)

    catlogHandler := handlers.CatalogHandler{}
	app.GET("/catalog", catlogHandler.HandleShowCataLog)

	app.Static("/css", "static/css")
	app.Static("/js", "static/js")
	app.Static("/vendors", "static/vendors")

    fmt.Println("Server is running on http://localhost:8080")
    app.Logger.Fatal(app.Start(":8080"))
}