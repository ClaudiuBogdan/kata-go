package articulationpoints

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindArticulationPoints(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     []int
	}{
		{
			name:     "Simple graph with one articulation point",
			vertices: 5,
			edges:    [][2]int{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}},
			want:     []int{0, 3},
		},
		{
			name:     "Graph with no articulation points",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     []int{},
		},
		{
			name:     "Graph with multiple articulation points",
			vertices: 7,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {1, 4}, {1, 6}, {3, 5}, {4, 5}},
			want:     []int{1, 5},
		},
		{
			name:     "Disconnected graph",
			vertices: 7,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}},
			want:     []int{},
		},
		{
			name:     "Linear graph",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:     []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindArticulationPoints()
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindArticulationPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	expectedAdjList := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2},
	}

	if !reflect.DeepEqual(g.AdjList, expectedAdjList) {
		t.Errorf("AddEdge() did not create the expected adjacency list. Got %v, want %v", g.AdjList, expectedAdjList)
	}
}
