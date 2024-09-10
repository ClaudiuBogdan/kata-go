package scc

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSCCs(t *testing.T) {
	tests := []struct {
		name     string
		graph    *Graph
		expected [][]int
	}{
		{
			name:     "Simple graph",
			graph:    createSimpleGraph(),
			expected: [][]int{{0, 1, 2}, {3}, {4}},
		},
		{
			name:     "Complex graph",
			graph:    createComplexGraph(),
			expected: [][]int{{0, 1, 2}, {3, 4}, {5, 6, 7}},
		},
		{
			name:     "Single vertex graph",
			graph:    createSingleVertexGraph(),
			expected: [][]int{{0}},
		},
		{
			name:     "Disconnected graph",
			graph:    createDisconnectedGraph(),
			expected: [][]int{{0}, {1}, {2}, {3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sccs := tt.graph.FindSCCs()
			if !equalSCCs(sccs, tt.expected) {
				t.Errorf("Expected SCCs %v, but got %v", tt.expected, sccs)
			}
		})
	}
}

func createSimpleGraph() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	return g
}

func createComplexGraph() *Graph {
	g := NewGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 5)
	return g
}

func createSingleVertexGraph() *Graph {
	return NewGraph(1)
}

func createDisconnectedGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 0)
	g.AddEdge(1, 1)
	g.AddEdge(2, 2)
	g.AddEdge(3, 3)
	return g
}

func equalSCCs(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each SCC and the slice of SCCs
	sortSCCs := func(sccs [][]int) {
		for _, scc := range sccs {
			sort.Ints(scc)
		}
		sort.Slice(sccs, func(i, j int) bool {
			return len(sccs[i]) > len(sccs[j]) || (len(sccs[i]) == len(sccs[j]) && less(sccs[i], sccs[j]))
		})
	}

	sortSCCs(a)
	sortSCCs(b)

	return reflect.DeepEqual(a, b)
}

func less(a, b []int) bool {
	for i := range a {
		if i >= len(b) {
			return false
		}
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
