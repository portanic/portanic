package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/db"
	"github.com/portanic/portanic/internal/handlers"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func connectZK() *zk.Conn {
	conn, _, err := zk.Connect([]string{"zookeeper:2181"}, time.Second*10)
	if err != nil {
		panic(err)
	}
	return conn
}

func discoverServices(conn *zk.Conn, path string) {
	children, _, err := conn.Children(path)
	if err != nil {
		fmt.Println("Failed to discover services:", err)
		return
	}
	for _, child := range children {
		fmt.Println("Discovered service:", child)
		data, _, _ := conn.Get(path + "/" + child)
		fmt.Println("Service data:", string(data))
	}
}

var (
	pluginPaths = []string{"./.plugins/", "/usr/share/portanic/plugins/"}
)

func main() {
	app := echo.New()

	zkConn := connectZK()
	defer zkConn.Close()

	// Discover services
	discoverServices(zkConn, "/services")

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

	catlogHandler := handlers.CatalogHandler{}
	app.GET("/catalog", catlogHandler.HandleShowCataLog)

	templatesHandler := handlers.TemplatesHandler{DB: queries}
	app.GET("/templates", templatesHandler.HandleShowTemplates)

	// Add route for creating a new template
	app.GET("/templates/new", templatesHandler.HandleNewTemplate)

	app.POST("/templates", templatesHandler.HandleCreateTemplate)

	app.Static("/css", "static/css")
	app.Static("/js", "static/js")
	app.Static("/vendors", "static/vendors")

	fmt.Println("Server is running on http://localhost:8080")
	app.Logger.Fatal(app.Start(":8080"))
}
