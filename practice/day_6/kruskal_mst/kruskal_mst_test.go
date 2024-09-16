package kruskalmst

import (
	"reflect"
	"sort"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	tests := []struct {
		name     string
		graph    *Graph
		expected []Edge
	}{
		{
			name:  "Simple graph",
			graph: createSimpleGraph(),
			expected: []Edge{
				{0, 1, 1},
				{1, 2, 2},
				{0, 3, 3},
			},
		},
		{
			name:  "Complex graph",
			graph: createComplexGraph(),
			expected: []Edge{
				{2, 3, 1},
				{0, 1, 2},
				{0, 3, 3},
				{1, 4, 4},
			},
		},
		{
			name:     "Single vertex graph",
			graph:    createSingleVertexGraph(),
			expected: []Edge{},
		},
		{
			name:  "Disconnected graph",
			graph: createDisconnectedGraph(),
			expected: []Edge{
				{0, 1, 1},
				{2, 3, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.graph.KruskalMST()
			if !equalEdgeSlices(result, tt.expected) {
				t.Errorf("KruskalMST() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func createSimpleGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(0, 3, 3)
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 5)
	return g
}

func createComplexGraph() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 3)
	g.AddEdge(1, 2, 5)
	g.AddEdge(1, 3, 6)
	g.AddEdge(1, 4, 4)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 8)
	return g
}

func createSingleVertexGraph() *Graph {
	return NewGraph(1)
}

func createDisconnectedGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 3, 2)
	return g
}

func equalEdgeSlices(a, b []Edge) bool {
	if len(a) != len(b) {
		return false
	}

	sortEdges := func(edges []Edge) {
		sort.Slice(edges, func(i, j int) bool {
			if edges[i].Src != edges[j].Src {
				return edges[i].Src < edges[j].Src
			}
			return edges[i].Dest < edges[j].Dest
		})
	}

	sortEdges(a)
	sortEdges(b)

	return reflect.DeepEqual(a, b)
}
