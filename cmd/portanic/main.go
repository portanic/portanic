package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", serveHome)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
    http.Handle("/vendors/", http.StripPrefix("/vendors/", http.FileServer(http.Dir("static/vendors/"))))
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}