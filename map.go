package main

type GridCell struct {
	Type     string `json:"type"`
	Contents string `json:"contents"`
}

type HexMap struct {
	Title   string       `json:"title"`
	Version uint64       `json:"version"`
	HexGrid [][]GridCell `json:"cells"`
}
