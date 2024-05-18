package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/db"
	"github.com/portanic/portanic/internal/handlers"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func main() {
	app := echo.New()

	connString := "postgres://portanic:portanic@db:5432/portanic"
	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pgxConnPool, err := pgxpool.NewWithConfig(context.TODO(), pgxConfig)
	if err != nil {
		panic(err)
	}

	defer pgxConnPool.Close()

	queries := db.New(pgxConnPool)

	homeHandler := handlers.HomeHandler{}
	app.GET("/", homeHandler.HandleHome)

	catlogHandler := handlers.CatalogHandler{DB: queries}
	app.GET("/catalog", catlogHandler.HandleShowCataLog)
	app.GET("/catalog/new", catlogHandler.HandleNewCatalog)
	app.POST("/catalog", catlogHandler.HandleCreateCatalog)

	templatesHandler := handlers.TemplatesHandler{DB: queries}
	app.GET("/templates", templatesHandler.HandleShowTemplates)
	app.GET("/templates/get", templatesHandler.HandleGetTemplateFields)
	app.GET("/templates/new", templatesHandler.HandleNewTemplate)
	app.POST("/templates", templatesHandler.HandleCreateTemplate)

	entryHandler := handlers.EntryHandler{DB: queries}
	app.GET("/entry/new", entryHandler.HandleNewEntry)
	app.POST("/entry", entryHandler.HandleCreateEntry)

	app.Static("/css", "static/css")
	app.Static("/js", "static/js")
	app.Static("/vendors", "static/vendors")

	fmt.Println("Server is running on http://localhost:8080")
	app.Logger.Fatal(app.Start(":8080"))
}
