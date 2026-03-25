package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "strconv"
)

type ResponseCellUpdated struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	GridCell GridCell `json:"grid-cell"`
}

type ErrorResponse struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
}

type MapListResponse struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	Maps     []string `json:"maps"`
}

type MapResponse struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	Map      HexMap   `json:"map"`
}

func StartRestController(hexMaps map[string]HexMap, transformTasks chan MapTransformTask) {
	// Define a route and handler for serving static files at the root
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))

	// Define a route and handler for load endpoint
	http.HandleFunc("/maps", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")

		keys := make([]string, 0, len(hexMaps))
		for k := range hexMaps {
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
	http.HandleFunc("/maps/{name}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		mapName := r.PathValue("name")

		hexMap := hexMaps[mapName]

		hexMap, exists := hexMaps[mapName]
		if !exists {
			res := ErrorResponse {
				Message: fmt.Sprintf("Map '%s' does not exist.", mapName),
				Status: 404,
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		res := MapResponse {
			Message:  "Success",
			Status:   200,
			Map:      hexMap,
		}
		json.NewEncoder(w).Encode(res)
	})

	// Define a route and handler for save endpoint
	http.HandleFunc("/maps/{name}/update/{row}/{col}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)

		mapName := r.PathValue("name")
		row, rowErr := strconv.Atoi(r.PathValue("row"))
		col, colErr := strconv.Atoi(r.PathValue("col"))
		if rowErr != nil || colErr != nil {
			res := ErrorResponse {
				Message: fmt.Sprintf("Map cell (%s, %s) does not exist.", r.PathValue("row"), r.PathValue("col")),
				Status: 404,
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		var newCell GridCell
		err := json.NewDecoder(r.Body).Decode(&newCell)
		if err != nil {
			res := ErrorResponse {
				Message: fmt.Sprintf("Bad request. Bad JSON body submitted.", mapName),
				Status: 400,
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Printf("  - Updating grid cell (%d, %d)) {type: '%s', contents: '%s'}\n", row, col, newCell.Type, newCell.Contents)

		w.Header().Set("Content-Type", "application/json")

		res := ResponseCellUpdated{
			Message:  "Accepted",
			Status:   202,
			GridCell: newCell,
		}

		hexMaps[mapName].HexGrid[row][col] = newCell

		json.NewEncoder(w).Encode(res)
	})

	// Start the server on port 8080
	log.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
