package bipartitematching

import (
	"testing"
)

func TestFindMaximumMatching(t *testing.T) {
	tests := []struct {
		name           string
		graph          *BipartiteGraph
		expectedSize   int
		expectedResult []int
	}{
		{
			name:           "Simple bipartite graph",
			graph:          createSimpleBipartiteGraph(),
			expectedSize:   2,
			expectedResult: []int{0, 1, -1},
		},
		{
			name:           "Complete bipartite graph",
			graph:          createCompleteBipartiteGraph(),
			expectedSize:   3,
			expectedResult: []int{0, 1, 2},
		},
		{
			name:           "Empty graph",
			graph:          createEmptyGraph(),
			expectedSize:   0,
			expectedResult: []int{},
		},
		{
			name:           "Unbalanced bipartite graph",
			graph:          createUnbalancedBipartiteGraph(),
			expectedSize:   2,
			expectedResult: []int{0, 1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.graph.FindMaximumMatching()
			size := tt.graph.GetMatchingSize(result)

			if size != tt.expectedSize {
				t.Errorf("Expected matching size %d, but got %d", tt.expectedSize, size)
			}

			if !isValidMatching(tt.graph, result) {
				t.Errorf("Invalid matching: %v", result)
			}

			if len(result) != len(tt.expectedResult) {
				t.Errorf("Expected result length %d, but got %d", len(tt.expectedResult), len(result))
			} else {
				for i := range result {
					if (result[i] == -1) != (tt.expectedResult[i] == -1) {
						t.Errorf("Mismatch at index %d: expected %d, got %d", i, tt.expectedResult[i], result[i])
					}
				}
			}
		})
	}
}

func createSimpleBipartiteGraph() *BipartiteGraph {
	g := NewBipartiteGraph(2, 3)
	g.AddEdge(0, 0)
	g.AddEdge(0, 1)
	g.AddEdge(1, 1)
	g.AddEdge(1, 2)
	return g
}

func createCompleteBipartiteGraph() *BipartiteGraph {
	g := NewBipartiteGraph(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.AddEdge(i, j)
		}
	}
	return g
}

func createEmptyGraph() *BipartiteGraph {
	return NewBipartiteGraph(0, 0)
}

func createUnbalancedBipartiteGraph() *BipartiteGraph {
	g := NewBipartiteGraph(2, 4)
	g.AddEdge(0, 0)
	g.AddEdge(0, 1)
	g.AddEdge(1, 1)
	g.AddEdge(1, 2)
	return g
}

func isValidMatching(g *BipartiteGraph, matching []int) bool {
	used := make([]bool, g.U)
	for v, u := range matching {
		if u != -1 {
			if used[u] {
				return false
			}
			used[u] = true

			found := false
			for _, edge := range g.Edges[u] {
				if edge == v {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
