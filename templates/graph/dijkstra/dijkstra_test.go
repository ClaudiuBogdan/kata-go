package dijkstra

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name        string
		vertices    int
		edges       [][3]int
		source      int
		wantDist    []int
		wantPrev    []int
	}{
		{
			name:     "Simple graph",
			vertices: 5,
			edges: [][3]int{
				{0, 1, 4}, {0, 2, 1}, {1, 3, 1}, {2, 1, 2},
				{2, 3, 5}, {3, 4, 3},
			},
			source:   0,
			wantDist: []int{0, 3, 1, 4, 7},
			wantPrev: []int{-1, 2, 0, 1, 3},
		},
		{
			name:     "Disconnected graph",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 2},
			},
			source:   0,
			wantDist: []int{0, 1, 3, 2147483647},
			wantPrev: []int{-1, 0, 1, -1},
		},
		{
			name:     "Graph with no edges",
			vertices: 3,
			edges:    [][3]int{},
			source:   0,
			wantDist: []int{0, 2147483647, 2147483647},
			wantPrev: []int{-1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			gotDist, gotPrev := g.Dijkstra(tt.source)

			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("Dijkstra() gotDist = %v, want %v", gotDist, tt.wantDist)
			}
			if !reflect.DeepEqual(gotPrev, tt.wantPrev) {
				t.Errorf("Dijkstra() gotPrev = %v, want %v", gotPrev, tt.wantPrev)
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
		want        []int
	}{
		{
			name:     "Path exists",
			vertices: 5,
			edges: [][3]int{
				{0, 1, 4}, {0, 2, 1}, {1, 3, 1}, {2, 1, 2},
				{2, 3, 5}, {3, 4, 3},
			},
			source:      0,
			destination: 4,
			want:        []int{0, 2, 1, 3, 4},
		},
		{
			name:     "No path",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 2},
			},
			source:      0,
			destination: 3,
			want:        nil,
		},
		{
			name:     "Source is destination",
			vertices: 3,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 2},
			},
			source:      1,
			destination: 1,
			want:        []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			got := g.GetPath(tt.source, tt.destination)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
