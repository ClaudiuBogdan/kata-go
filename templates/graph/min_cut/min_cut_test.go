package mincut

import (
	"sort"
	"testing"
)

func TestFindMinCut(t *testing.T) {
	tests := []struct {
		name            string
		graph           *Graph
		expectedMinCut  int
		expectedPartition []int
	}{
		{
			name:            "Simple graph",
			graph:           createSimpleGraph(),
			expectedMinCut:  2,
			expectedPartition: []int{2, 3},
		},
		{
			name:            "Complex graph",
			graph:           createComplexGraph(),
			expectedMinCut:  4,
			expectedPartition: []int{0, 1},
		},
		{
			name:            "Disconnected graph",
			graph:           createDisconnectedGraph(),
			expectedMinCut:  0,
			expectedPartition: []int{2, 3},
		},
		{
			name:            "Complete graph",
			graph:           createCompleteGraph(4),
			expectedMinCut:  3,
			expectedPartition: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minCut, partition := tt.graph.FindMinCut()
			if minCut != tt.expectedMinCut {
				t.Errorf("Expected min cut %d, but got %d", tt.expectedMinCut, minCut)
			}

			if !equalPartitions(partition, tt.expectedPartition) {
				t.Errorf("Expected partition %v, but got %v", tt.expectedPartition, partition)
			}

			if !isValidCut(tt.graph, partition, minCut) {
				t.Errorf("Invalid cut: partition %v with cut value %d", partition, minCut)
			}
		})
	}
}

func createSimpleGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 2, 3)
	g.AddEdge(0, 3, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	return g
}

func createComplexGraph() *Graph {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 3)
	g.AddEdge(1, 4, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 4)
	return g
}

func createDisconnectedGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 3, 1)
	return g
}

func createCompleteGraph(n int) *Graph {
	g := NewGraph(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.AddEdge(i, j, 1)
		}
	}
	return g
}

func equalPartitions(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := make([]int, len(a))
	sortedB := make([]int, len(b))
	copy(sortedA, a)
	copy(sortedB, b)
	sort.Ints(sortedA)
	sort.Ints(sortedB)
	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}
	return true
}

func isValidCut(g *Graph, partition []int, cutValue int) bool {
	inPartition := make(map[int]bool)
	for _, v := range partition {
		inPartition[v] = true
	}

	actualCut := 0
	for i := 0; i < g.V; i++ {
		for j := i + 1; j < g.V; j++ {
			if inPartition[i] != inPartition[j] {
				actualCut += g.Edges[i][j]
			}
		}
	}

	return actualCut == cutValue
}
