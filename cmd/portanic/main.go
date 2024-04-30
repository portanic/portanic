package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	//"github.com/portanic/portanic/internal/db"
	"github.com/go-zookeeper/zk"
	"github.com/portanic/portanic/internal/handlers"
	//"github.com/portanic/portanic/plugins"
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

	//config := db.Config{
	//Host:     "db",
	//Port:     5432,
	//User:     "portanic",
	//Password: "portanic",
	//DBName:   "portanic",
	//}

	//database, err := db.New(config)
	//if err != nil {
	//log.Fatalf("Failed to connect to database: %v", err)
	//}
	//defer database.Close()

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
