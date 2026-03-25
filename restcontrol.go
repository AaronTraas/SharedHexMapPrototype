package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	GridCell GridCell `json:"grid-cell"`
}

type MapListResponse struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	Maps     []string   `json:"maps"`
}

var cell GridCell

func handleLoadMapCell(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
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
}

func handleSaveMapCell(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
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

	fmt.Printf("  - Updating grid cell {type: '%s', contents: '%s'\n", newCell.Type, newCell.Contents)

	cell.Type = newCell.Type
	cell.Contents = newCell.Contents

	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Message:  "Accepted",
		Status:   202,
		GridCell: cell,
	}

	json.NewEncoder(w).Encode(res)
}

func StartRestController(hexMaps *map[string]HexMap, transformTasks chan MapTransformTask) {
	// Define a route and handler for serving static files at the root
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))

	cell = GridCell{
		Type:     "default",
		Contents: "Hello World",
	}

	// Define a route and handler for load endpoint
	http.HandleFunc("/maps", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		keys := make([]string, 0, len(*hexMaps))
		for k := range *hexMaps {
			keys = append(keys, k)
		}

		res := MapListResponse {
			Message:  "Success",
			Status:   200,
			Maps:     keys,
		}
		json.NewEncoder(w).Encode(res)
	})

	// Define a route and handler for load endpoint
	http.HandleFunc("/api/load", handleLoadMapCell)

	// Define a route and handler for save endpoint
	http.HandleFunc("/api/save", handleSaveMapCell)

	// Start the server on port 8080
	fmt.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
