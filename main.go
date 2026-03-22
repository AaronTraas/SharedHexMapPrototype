package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type GridCell struct {
  Id       string `json:"id"`
  Contents string `json:"contents"`
}

type Response struct {
  Message  string   `json:"message"`
  Status   int      `json:"status"`
  GridCell GridCell `json:"grid-cell"`
}

func main() {
  // Define a route and handler for serving static files at the root
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", http.StripPrefix("/", fs))

  // Define a route and handler for load endpoint
  http.HandleFunc("/api/load", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    cell := GridCell{
      Id: "default",
      Contents: "Hello World",
    }

    res := Response{
      Message: "Success",
      Status: 200,
      GridCell: cell,
    }

    json.NewEncoder(w).Encode(res)
  })

  // Define a route and handler for save endpoint
  http.HandleFunc("/api/save", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    res := Response{
      Message: "Success",
      Status: 200,
    }

    json.NewEncoder(w).Encode(res)
  })

  // Start the server on port 8080
  fmt.Println("Server starting on http://localhost:8080...")
  http.ListenAndServe(":8080", nil)
}
