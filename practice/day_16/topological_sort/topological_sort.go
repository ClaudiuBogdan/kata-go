package toposort

import "fmt"

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
	order := make([]int, 0)
	state := make(map[int]int)

	const (
		visiting = 1
		visited  = 2
	)
	var dfs func(v int) error

	dfs = func(v int) error {
		if state[v] == visited {
			return nil
		}
		if state[v] == visiting {
			return fmt.Errorf("graph contains a cycle")
		}

		state[v] = visiting

		for _, neighbor := range g.Edges[v] {
			if err := dfs(neighbor); err != nil {
				return err
			}
		}

		state[v] = visited
		order = append([]int{v}, order...)
		return nil
	}

	for v := 0; v < g.V; v++ {
		if err := dfs(v); err != nil {
			return nil, err
		}
	}

	return order, nil
}
