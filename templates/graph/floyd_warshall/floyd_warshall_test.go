package floydwarshall

import (
	"math"
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][3]float64
		want     [][]float64
	}{
		{
			name:     "Simple graph",
			vertices: 4,
			edges: [][3]float64{
				{0, 1, 5},
				{0, 3, 10},
				{1, 2, 3},
				{2, 3, 1},
			},
			want: [][]float64{
				{0, 5, 8, 9},
				{math.Inf(1), 0, 3, 4},
				{math.Inf(1), math.Inf(1), 0, 1},
				{math.Inf(1), math.Inf(1), math.Inf(1), 0},
			},
		},
		{
			name:     "Graph with negative edges",
			vertices: 3,
			edges: [][3]float64{
				{0, 1, 3},
				{1, 2, -2},
				{0, 2, 5},
			},
			want: [][]float64{
				{0, 3, 1},
				{math.Inf(1), 0, -2},
				{math.Inf(1), math.Inf(1), 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(int(edge[0]), int(edge[1]), edge[2])
			}

			got := g.FloydWarshall()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloydWarshall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPath(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][3]float64
		from     int
		to       int
		want     []int
	}{
		{
			name:     "Path exists",
			vertices: 4,
			edges: [][3]float64{
				{0, 1, 5},
				{0, 3, 10},
				{1, 2, 3},
				{2, 3, 1},
			},
			from: 0,
			to:   3,
			want: []int{0, 1, 2, 3},
		},
		{
			name:     "No path",
			vertices: 3,
			edges: [][3]float64{
				{0, 1, 5},
				{1, 2, 3},
			},
			from: 2,
			to:   0,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(int(edge[0]), int(edge[1]), edge[2])
			}

			got := g.GetPath(tt.from, tt.to)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasNegativeCycle(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][3]float64
		want     bool
	}{
		{
			name:     "No negative cycle",
			vertices: 3,
			edges: [][3]float64{
				{0, 1, 1},
				{1, 2, -1},
				{2, 0, 1},
			},
			want: false,
		},
		{
			name:     "With negative cycle",
			vertices: 3,
			edges: [][3]float64{
				{0, 1, 1},
				{1, 2, -1},
				{2, 0, -1},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(int(edge[0]), int(edge[1]), edge[2])
			}

			got := g.HasNegativeCycle()

			if got != tt.want {
				t.Errorf("HasNegativeCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
