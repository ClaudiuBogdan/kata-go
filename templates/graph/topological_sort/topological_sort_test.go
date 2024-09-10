package toposort

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name           string
		graph          *Graph
		expectedOrder  []int
		expectSuccess  bool
	}{
		{
			name:           "Simple DAG",
			graph:          createSimpleDAG(),
			expectedOrder:  []int{0, 1, 2, 3},
			expectSuccess:  true,
		},
		{
			name:           "Complex DAG",
			graph:          createComplexDAG(),
			expectedOrder:  []int{5, 4, 2, 3, 1, 0},
			expectSuccess:  true,
		},
		{
			name:           "Single vertex graph",
			graph:          createSingleVertexGraph(),
			expectedOrder:  []int{0},
			expectSuccess:  true,
		},
		{
			name:           "Graph with cycle",
			graph:          createGraphWithCycle(),
			expectedOrder:  nil,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, success := tt.graph.TopologicalSort()
			if success != tt.expectSuccess {
				t.Errorf("Expected success: %v, got: %v", tt.expectSuccess, success)
			}
			if !reflect.DeepEqual(order, tt.expectedOrder) {
				t.Errorf("Expected order %v, but got %v", tt.expectedOrder, order)
			}
			if success {
				if !isValidTopologicalOrder(tt.graph, order) {
					t.Errorf("Invalid topological order: %v", order)
				}
			}
		})
	}
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
