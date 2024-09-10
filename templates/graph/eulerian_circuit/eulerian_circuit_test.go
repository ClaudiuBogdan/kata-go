package euleriancircuit

import (
	"reflect"
	"testing"
)

func TestHasEulerianCircuit(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     bool
	}{
		{
			name:     "Simple Eulerian circuit",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     true,
		},
		{
			name:     "Non-Eulerian graph (odd degree)",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}},
			want:     false,
		},
		{
			name:     "Disconnected graph",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {2, 3}},
			want:     false,
		},
		{
			name:     "Empty graph",
			vertices: 3,
			edges:    [][2]int{},
			want:     true,
		},
		{
			name:     "Complex Eulerian circuit",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}, {0, 2}, {2, 4}, {4, 1}, {1, 3}, {3, 0}},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := g.HasEulerianCircuit(); got != tt.want {
				t.Errorf("HasEulerianCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindEulerianCircuit(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     []int
	}{
		{
			name:     "Simple Eulerian circuit",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     []int{0, 1, 2, 0},
		},
		{
			name:     "Non-Eulerian graph",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}},
			want:     nil,
		},
		{
			name:     "Empty graph",
			vertices: 3,
			edges:    [][2]int{},
			want:     []int{0},
		},
		{
			name:     "Complex Eulerian circuit",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}, {0, 2}, {2, 4}, {4, 1}, {1, 3}, {3, 0}},
			want:     []int{0, 1, 2, 3, 4, 1, 3, 0, 2, 4, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindEulerianCircuit()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindEulerianCircuit() = %v, want %v", got, tt.want)
			}

			if got != nil && !isValidEulerianCircuit(g, got) {
				t.Errorf("FindEulerianCircuit() returned an invalid Eulerian circuit: %v", got)
			}
		})
	}
}

// isValidEulerianCircuit checks if the given circuit is a valid Eulerian circuit
func isValidEulerianCircuit(g *Graph, circuit []int) bool {
	if len(circuit) == 0 || circuit[0] != circuit[len(circuit)-1] {
		return false
	}

	edgeCount := make(map[[2]int]int)
	for v, edges := range g.AdjList {
		for _, u := range edges {
			if v < u {
				edgeCount[[2]int{v, u}]++
			} else {
				edgeCount[[2]int{u, v}]++
			}
		}
	}

	for i := 0; i < len(circuit)-1; i++ {
		v, u := circuit[i], circuit[i+1]
		if v > u {
			v, u = u, v
		}
		edge := [2]int{v, u}
		if edgeCount[edge] == 0 {
			return false
		}
		edgeCount[edge]--
	}

	for _, count := range edgeCount {
		if count != 0 {
			return false
		}
	}

	return true
}
