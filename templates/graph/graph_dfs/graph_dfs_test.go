package dfs

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		source   int
		want     []int
	}{
		{
			name:     "Simple connected graph",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}},
			source:   0,
			want:     []int{0, 1, 2, 3},
		},
		{
			name:     "Disconnected graph",
			vertices: 6,
			edges:    [][2]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}},
			source:   0,
			want:     []int{0, 1, 2},
		},
		{
			name:     "Single vertex graph",
			vertices: 1,
			edges:    [][2]int{},
			source:   0,
			want:     []int{0},
		},
		{
			name:     "Linear graph",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			source:   0,
			want:     []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.DFS(tt.source)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPath(t *testing.T) {
	tests := []struct {
		name        string
		vertices    int
		edges       [][2]int
		source      int
		destination int
		want        bool
	}{
		{
			name:        "Path exists",
			vertices:    5,
			edges:       [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}},
			source:      0,
			destination: 4,
			want:        true,
		},
		{
			name:        "No path",
			vertices:    5,
			edges:       [][2]int{{0, 1}, {1, 2}, {3, 4}},
			source:      0,
			destination: 4,
			want:        false,
		},
		{
			name:        "Source is destination",
			vertices:    3,
			edges:       [][2]int{{0, 1}, {1, 2}},
			source:      1,
			destination: 1,
			want:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.HasPath(tt.source, tt.destination)

			if got != tt.want {
				t.Errorf("HasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllPaths(t *testing.T) {
	tests := []struct {
		name        string
		vertices    int
		edges       [][2]int
		source      int
		destination int
		want        [][]int
	}{
		{
			name:        "Multiple paths",
			vertices:    5,
			edges:       [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}},
			source:      0,
			destination: 4,
			want:        [][]int{{0, 1, 3, 4}, {0, 2, 3, 4}},
		},
		{
			name:        "No paths",
			vertices:    5,
			edges:       [][2]int{{0, 1}, {1, 2}, {3, 4}},
			source:      0,
			destination: 4,
			want:        [][]int{},
		},
		{
			name:        "Single path",
			vertices:    4,
			edges:       [][2]int{{0, 1}, {1, 2}, {2, 3}},
			source:      0,
			destination: 3,
			want:        [][]int{{0, 1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindAllPaths(tt.source, tt.destination)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
