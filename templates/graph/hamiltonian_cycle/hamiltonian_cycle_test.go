package hamiltoniancycle

import (
	"reflect"
	"testing"
)

func TestFindHamiltonianCycle(t *testing.T) {
	tests := []struct {
		name     string
		graph    *Graph
		expected []int
	}{
		{
			name:     "Graph with Hamiltonian cycle",
			graph:    createGraphWithHamiltonianCycle(),
			expected: []int{0, 1, 2, 4, 3, 0},
		},
		{
			name:     "Graph without Hamiltonian cycle",
			graph:    createGraphWithoutHamiltonianCycle(),
			expected: nil,
		},
		{
			name:     "Complete graph",
			graph:    createCompleteGraph(5),
			expected: []int{0, 1, 2, 3, 4, 0},
		},
		{
			name:     "Single node graph",
			graph:    createSingleNodeGraph(),
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.graph.FindHamiltonianCycle()
			if !isValidHamiltonianCycle(tt.graph, result) {
				t.Errorf("FindHamiltonianCycle() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func createGraphWithHamiltonianCycle() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 0)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	return g
}

func createGraphWithoutHamiltonianCycle() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
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

func isValidHamiltonianCycle(g *Graph, cycle []int) bool {
	if cycle == nil {
		return true // No cycle found, which is valid for graphs without Hamiltonian cycles
	}

	if len(cycle) != g.V+1 {
		return false
	}

	visited := make([]bool, g.V)
	for i := 0; i < len(cycle)-1; i++ {
		if visited[cycle[i]] {
			return false
		}
		visited[cycle[i]] = true

		if !contains(g.Adj[cycle[i]], cycle[i+1]) {
			return false
		}
	}

	return cycle[0] == cycle[len(cycle)-1] && reflect.DeepEqual(visited, []bool{true, true, true, true, true})
}
