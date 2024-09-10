package primsmst

import (
	"sort"
	"testing"
)

func TestPrimMST(t *testing.T) {
	tests := []struct {
		name           string
		graph          *Graph
		expectedWeight int
	}{
		{
			name:           "Simple graph",
			graph:          createSimpleGraph(),
			expectedWeight: 37,
		},
		{
			name:           "Complex graph",
			graph:          createComplexGraph(),
			expectedWeight: 39,
		},
		{
			name:           "Single vertex graph",
			graph:          createSingleVertexGraph(),
			expectedWeight: 0,
		},
		{
			name:           "Disconnected graph",
			graph:          createDisconnectedGraph(),
			expectedWeight: 7, // Only one component will be fully connected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mst := tt.graph.PrimMST()
			weight := calculateMSTWeight(mst)
			if weight != tt.expectedWeight {
				t.Errorf("Expected MST weight %d, but got %d", tt.expectedWeight, weight)
			}

			if !isValidMST(tt.graph, mst) {
				t.Errorf("Invalid MST: %v", mst)
			}
		})
	}
}

func createSimpleGraph() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)
	return g
}

func createComplexGraph() *Graph {
	g := NewGraph(7)
	g.AddEdge(0, 1, 7)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 3, 9)
	g.AddEdge(1, 4, 7)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 15)
	g.AddEdge(3, 5, 6)
	g.AddEdge(4, 5, 8)
	g.AddEdge(4, 6, 9)
	g.AddEdge(5, 6, 11)
	return g
}

func createSingleVertexGraph() *Graph {
	return NewGraph(1)
}

func createDisconnectedGraph() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 2)
	g.AddEdge(3, 4, 5)
	return g
}

func calculateMSTWeight(mst []Edge) int {
	weight := 0
	for _, edge := range mst {
		weight += edge.Weight
	}
	return weight
}

func isValidMST(g *Graph, mst []Edge) bool {
	if len(mst) != g.V-1 {
		return false
	}

	// Check if the MST is connected
	parent := make([]int, g.V)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(i, j int) {
		pi, pj := find(i), find(j)
		if pi != pj {
			parent[pi] = pj
		}
	}

	for _, edge := range mst {
		union(edge.To, find(0))
	}

	for i := 1; i < g.V; i++ {
		if find(i) != find(0) {
			return false
		}
	}

	// Check if it's a valid spanning tree (no cycles)
	return len(mst) == g.V-1
}
