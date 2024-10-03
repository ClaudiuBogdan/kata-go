package toposort

import (
	"math/rand"
	"testing"
	"time"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph
		expectSuccess bool
	}{
		{
			name:          "Simple DAG",
			graph:         createSimpleDAG(),
			expectSuccess: true,
		},
		{
			name:          "Complex DAG",
			graph:         createComplexDAG(),
			expectSuccess: true,
		},
		{
			name:          "Single vertex graph",
			graph:         createSingleVertexGraph(),
			expectSuccess: true,
		},
		{
			name:          "Graph with cycle",
			graph:         createGraphWithCycle(),
			expectSuccess: false,
		},
		{
			name:          "Empty graph",
			graph:         NewGraph(0),
			expectSuccess: true,
		},
		{
			name:          "Disconnected graph",
			graph:         createDisconnectedGraph(),
			expectSuccess: true,
		},
		{
			name:          "Graph with multiple cycles",
			graph:         createGraphWithMultipleCycles(),
			expectSuccess: false,
		},
		{
			name:          "Graph with no edges",
			graph:         createGraphWithNoEdges(5),
			expectSuccess: true,
		},
		{
			name:          "Long chain graph",
			graph:         createLongChainGraph(10),
			expectSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := tt.graph.TopologicalSort()
			if (err == nil) != tt.expectSuccess {
				t.Errorf("Expected success: %v, got error: %v", tt.expectSuccess, err)
			}
			if tt.expectSuccess {
				if !isValidTopologicalOrder(tt.graph, order) {
					t.Errorf("Invalid topological order: %v", order)
				}
			} else {
				if err == nil || err.Error() != "graph contains a cycle" {
					t.Errorf("Expected 'graph contains a cycle' error, got: %v", err)
				}
			}
		})
	}
}

func TestMultipleValidOrderings(t *testing.T) {
	g := createGraphWithMultipleValidOrderings()
	order, err := g.TopologicalSort()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !isValidTopologicalOrder(g, order) {
		t.Errorf("Invalid topological order: %v", order)
	}
	// Note: We don't check for a specific ordering, just that it's valid
}

func TestRandomDAGs(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		g := createRandomDAG(20, 0.2)
		order, err := g.TopologicalSort()
		if err != nil {
			t.Errorf("Error in random DAG #%d: %v", i, err)
		} else if !isValidTopologicalOrder(g, order) {
			t.Errorf("Invalid topological order for random DAG #%d: %v", i, order)
		}
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	g := createLargeDAG(1000, 0.01)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := g.TopologicalSort()
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func createRandomDAG(n int, edgeProbability float64) *Graph {
	g := NewGraph(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rand.Float64() < edgeProbability {
				g.AddEdge(i, j)
			}
		}
	}
	return g
}

func createLargeDAG(n int, edgeProbability float64) *Graph {
	return createRandomDAG(n, edgeProbability)
}

func isValidTopologicalOrder(g *Graph, order []int) bool {
	if len(order) != g.V {
		return false
	}

	position := make([]int, g.V)
	for i, v := range order {
		position[v] = i
	}

	for v := 0; v < g.V; v++ {
		for _, w := range g.Edges[v] {
			if position[v] > position[w] {
				return false
			}
		}
	}

	return true
}


func createSimpleDAG() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	return g
}

func createComplexDAG() *Graph {
	g := NewGraph(6)
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	return g
}

func createSingleVertexGraph() *Graph {
	return NewGraph(1)
}

func createGraphWithCycle() *Graph {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	return g
}

func createDisconnectedGraph() *Graph {
	g := NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)
	return g
}

func createGraphWithMultipleCycles() *Graph {
	g := NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 3)
	return g
}

func createGraphWithNoEdges(n int) *Graph {
	return NewGraph(n)
}

func createLongChainGraph(n int) *Graph {
	g := NewGraph(n)
	for i := 0; i < n-1; i++ {
		g.AddEdge(i, i+1)
	}
	return g
}

func createGraphWithMultipleValidOrderings() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	return g
}