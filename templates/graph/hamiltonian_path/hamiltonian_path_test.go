package hamiltonianpath

import (
	"testing"
)

func TestFindHamiltonianPath(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph
		expectedExist bool
	}{
		{
			name:          "Graph with Hamiltonian path",
			graph:         createGraphWithHamiltonianPath(),
			expectedExist: true,
		},
		{
			name:          "Graph without Hamiltonian path",
			graph:         createGraphWithoutHamiltonianPath(),
			expectedExist: false,
		},
		{
			name:          "Complete graph",
			graph:         createCompleteGraph(5),
			expectedExist: true,
		},
		{
			name:          "Single node graph",
			graph:         createSingleNodeGraph(),
			expectedExist: true,
		},
		{
			name:          "Two node graph",
			graph:         createTwoNodeGraph(),
			expectedExist: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.graph.FindHamiltonianPath()
			if tt.expectedExist && !isValidHamiltonianPath(tt.graph, result) {
				t.Errorf("FindHamiltonianPath() = %v, expected a valid Hamiltonian path", result)
			}
			if !tt.expectedExist && result != nil {
				t.Errorf("FindHamiltonianPath() = %v, expected no Hamiltonian path", result)
			}
		})
	}
}

func createGraphWithHamiltonianPath() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	return g
}

func createGraphWithoutHamiltonianPath() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	return g
}

func createCompleteGraph(n int) *Graph {
	g := NewGraph(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.AddEdge(i, j)
		}
	}
	return g
}

func createSingleNodeGraph() *Graph {
	return NewGraph(1)
}

func createTwoNodeGraph() *Graph {
	g := NewGraph(2)
	g.AddEdge(0, 1)
	return g
}

func isValidHamiltonianPath(g *Graph, path []int) bool {
	if path == nil {
		return false
	}

	if len(path) != g.V {
		return false
	}

	visited := make([]bool, g.V)
	for i := 0; i < len(path)-1; i++ {
		if visited[path[i]] {
			return false
		}
		visited[path[i]] = true

		if !contains(g.Adj[path[i]], path[i+1]) {
			return false
		}
	}

	return !visited[path[len(path)-1]]
}
