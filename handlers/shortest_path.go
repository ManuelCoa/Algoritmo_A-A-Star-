package handlers

import (
	algorithms "astar-service/algoritmo"
	"astar-service/models"
	"encoding/json"
	"net/http"
)

type Request struct {
	Graph  models.Graph `json:"graph"`
	Start  string       `json:"start"`
	Target string       `json:"target"`
}

type Response struct {
	Path []string `json:"path"`
	Cost int      `json:"cost"`
}

func ShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	path, cost, err := algorithms.AStar(req.Graph, req.Start, req.Target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{Path: path, Cost: cost}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
