package bellmanford

import (
	"reflect"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	tests := []struct {
		name           string
		vertices       int
		edges          [][3]int
		source         int
		wantDist       []int
		wantNegCycle   bool
	}{
		{
			name:     "Simple graph",
			vertices: 5,
			edges: [][3]int{
				{0, 1, -1}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2},
				{1, 4, 2}, {3, 2, 5}, {3, 1, 1}, {4, 3, -3},
			},
			source:       0,
			wantDist:     []int{0, -1, 2, -2, 1},
			wantNegCycle: false,
		},
		{
			name:     "Graph with negative cycle",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, -1}, {2, 3, -1}, {3, 1, -1},
			},
			source:       0,
			wantDist:     nil,
			wantNegCycle: true,
		},
		{
			name:     "Disconnected graph",
			vertices: 5,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 2}, {3, 4, 3},
			},
			source:       0,
			wantDist:     []int{0, 1, 3, 2147483647, 2147483647},
			wantNegCycle: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			gotDist, gotNegCycle := g.BellmanFord(tt.source)

			if gotNegCycle != tt.wantNegCycle {
				t.Errorf("BellmanFord() gotNegCycle = %v, want %v", gotNegCycle, tt.wantNegCycle)
			}

			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("BellmanFord() gotDist = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}

func TestGetPath(t *testing.T) {
	tests := []struct {
		name        string
		vertices    int
		edges       [][3]int
		source      int
		destination int
		wantPath    []int
	}{
		{
			name:     "Simple path",
			vertices: 5,
			edges: [][3]int{
				{0, 1, -1}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2},
				{1, 4, 2}, {3, 2, 5}, {3, 1, 1}, {4, 3, -3},
			},
			source:      0,
			destination: 3,
			wantPath:    []int{0, 1, 4, 3},
		},
		{
			name:     "No path",
			vertices: 5,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 2}, {3, 4, 3},
			},
			source:      0,
			destination: 4,
			wantPath:    nil,
		},
		{
			name:     "Path with negative cycle",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, -1}, {2, 3, -1}, {3, 1, -1},
			},
			source:      0,
			destination: 3,
			wantPath:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			gotPath := g.GetPath(tt.source, tt.destination)

			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("GetPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
