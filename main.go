package main

import (
  "fmt"
  "net/http"
)

func main() {
  // Define a route and handler for serving static files at the root
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", http.StripPrefix("/", fs))

  // Define a route and handler for services
  http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Welcome to my Go web service!")
  })

  // Start the server on port 8080
  fmt.Println("Server starting on http://localhost:8080...")
  http.ListenAndServe(":8080", nil)
}
