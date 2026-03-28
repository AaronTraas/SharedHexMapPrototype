package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MapTransformTask struct {
	MapId     string
	Cell      GridCell
	Timestamp time.Time
}

var transformTaskChannel chan MapTransformTask

func mapTransformWorker() {
	for task := range transformTaskChannel {
		fmt.Printf("Worker processing task: %s\n", task.Timestamp)
		// Simulate some work
		// time.Sleep(time.Second)
		log.Printf("Worker finished job: %s\n", task.Timestamp)
	}
}

func loadMap(jsonFilePath string) (HexMap, error) {
	log.Printf("- Loading '%s'\n", jsonFilePath)
	var hexmap HexMap

	jsonFile, err := os.Open(jsonFilePath)
	// if os.Open returns an error then handle it
	if err != nil {
		return hexmap, err
	}

	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return hexmap, err
	}

	err = json.Unmarshal(byteValue, &hexmap)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return hexmap, err
	}

	return hexmap, nil
}

func main() {
	// Create in memory data store for all maps
	hexMaps := make(map[string]HexMap)

	dirPath := "./maps" // Current directory, replace with your path

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fullPath := filepath.Join(dirPath, file.Name())
			hexMap, err := loadMap(fullPath)
			if err != nil {
				log.Fatal(err)
			}

			hexMaps[file.Name()] = hexMap
		}
	}

	// Create go channel for map transform operations to be created by the web service
	// and consumed by the consumer routine
	transformTaskChannel := make(chan MapTransformTask)

	go mapTransformWorker()
	StartRestController(hexMaps, transformTaskChannel)
}
