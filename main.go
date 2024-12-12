package main

import (
	"astar-service/handlers"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/shortest-path", handlers.ShortestPathHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
