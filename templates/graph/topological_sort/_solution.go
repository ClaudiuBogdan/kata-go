package toposort

import (
	"fmt"
)

// Graph represents a directed graph using an adjacency list.
type Graph struct {
	V     int     // Number of vertices
	Edges [][]int // Adjacency list
}

// NewGraph creates a new graph with V vertices.
func NewGraph(V int) *Graph {
	return &Graph{
		V:     V,
		Edges: make([][]int, V),
	}
}

// AddEdge adds a directed edge from v to w.
func (g *Graph) AddEdge(v, w int) {
	g.Edges[v] = append(g.Edges[v], w)
}

// TopologicalSort performs a topological sort of the graph.
// It returns the sorted order and an error if a cycle is detected.
func (g *Graph) TopologicalSort() ([]int, error) {
	visited := make([]bool, g.V)
	stack := make([]bool, g.V)
	order := make([]int, 0, g.V)

	var dfs func(v int) error
	dfs = func(v int) error {
		visited[v] = true
		stack[v] = true

		for _, w := range g.Edges[v] {
			if !visited[w] {
				if err := dfs(w); err != nil {
					return err
				}
			} else if stack[w] {
				return fmt.Errorf("graph contains a cycle")
			}
		}

		stack[v] = false
		order = append([]int{v}, order...) // Prepend v to order
		return nil
	}

	for v := 0; v < g.V; v++ {
		if !visited[v] {
			if err := dfs(v); err != nil {
				return nil, err
			}
		}
	}

	return order, nil
}