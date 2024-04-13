package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", serveHome)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
}
