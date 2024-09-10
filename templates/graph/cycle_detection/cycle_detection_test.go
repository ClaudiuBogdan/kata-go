package cycledetection

import (
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     bool
	}{
		{
			name:     "Cyclic graph",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
			want:     true,
		},
		{
			name:     "Acyclic graph",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}},
			want:     false,
		},
		{
			name:     "Self-loop",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 2}},
			want:     true,
		},
		{
			name:     "Disconnected cyclic graph",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 0}, {2, 3}, {3, 4}},
			want:     true,
		},
		{
			name:     "Empty graph",
			vertices: 3,
			edges:    [][2]int{},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := g.DetectCycle(); got != tt.want {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCycle(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     []int
	}{
		{
			name:     "Simple cycle",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     []int{0, 1, 2, 0},
		},
		{
			name:     "No cycle",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}},
			want:     nil,
		},
		{
			name:     "Self-loop",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 2}},
			want:     []int{2, 2},
		},
		{
			name:     "Multiple cycles",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}, {3, 4}, {4, 2}},
			want:     []int{0, 1, 2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindCycle()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
