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

	cell := GridCell{
		Id:       "default",
		Contents: "Hello World",
	}

	// Define a route and handler for load endpoint
	http.HandleFunc("/api/load", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/load")
		if r.Method != http.MethodGet {
			fmt.Println("  - method not allowed")
			http.Error(w, "405: Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Println("  - Returning grid cell")
		w.Header().Set("Content-Type", "application/json")

		res := Response{
			Message:  "Success",
			Status:   200,
			GridCell: cell,
		}

		json.NewEncoder(w).Encode(res)
	})

	// Define a route and handler for save endpoint
	http.HandleFunc("/api/save", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/save")
		if r.Method != http.MethodPost {
			fmt.Println("  - method not allowed")
			http.Error(w, "405: Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var newCell GridCell
		err := json.NewDecoder(r.Body).Decode(&newCell)
		if err != nil {
			// return HTTP 400 bad request
			fmt.Println("  - Bad body")
			http.Error(w, "400: Bad request", http.StatusBadRequest)
			return
		}

		fmt.Printf("  - Updating grid cell {id: '%s', contents: '%s'\n", newCell.Id, newCell.Contents)

		cell.Id = newCell.Id
		cell.Contents = newCell.Contents

		w.Header().Set("Content-Type", "application/json")

		res := Response{
			Message:  "Accepted",
			Status:   202,
			GridCell: cell,
		}

		json.NewEncoder(w).Encode(res)
	})

	// Start the server on port 8080
	fmt.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
