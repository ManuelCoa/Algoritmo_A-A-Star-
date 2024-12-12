package models

type Graph struct {
	Nodes map[string]Node `json:"nodes"`
}

type Node struct {
	Edges     map[string]int `json:"edges"`
	Heuristic int            `json:"heuristic"` // Estimación de distancia al objetivo
}
