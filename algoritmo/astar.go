package algorithms

import (
	"astar-service/models"
	"container/heap"
	"errors"
)

type Item struct {
	Node     string
	Priority int
	Path     []string
	Cost     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func AStar(graph models.Graph, start, target string) ([]string, int, error) {
	if _, ok := graph.Nodes[start]; !ok {
		return nil, 0, errors.New("start node not found")
	}
	if _, ok := graph.Nodes[target]; !ok {
		return nil, 0, errors.New("target node not found")
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: start, Priority: 0, Path: []string{start}, Cost: 0})

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		if current.Node == target {
			return current.Path, current.Cost, nil
		}
		if visited[current.Node] {
			continue
		}
		visited[current.Node] = true

		for neighbor, cost := range graph.Nodes[current.Node].Edges {
			if visited[neighbor] {
				continue
			}
			newCost := current.Cost + cost
			priority := newCost + graph.Nodes[neighbor].Heuristic
			newPath := append([]string{}, current.Path...)
			newPath = append(newPath, neighbor)
			heap.Push(pq, &Item{Node: neighbor, Priority: priority, Path: newPath, Cost: newCost})
		}
	}

	return nil, 0, errors.New("no path found")
}
