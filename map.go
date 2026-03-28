package main

type GridCell string

type HexMap struct {
	Title   string               `json:"title"`
	Version uint64               `json:"version"`
	Default GridCell             `json:"default"`
	MinX int                     `json:"min_x"`
	MaxX int                     `json:"max_x"`
	MinY int                     `json:"min_y"`
	MaxY int                     `json:"max_y"`
	HexGrid map[string]GridCell  `json:"cells"`
}
