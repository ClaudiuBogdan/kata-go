package bridges

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindBridges(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     [][2]int
	}{
		{
			name:     "Simple graph with one bridge",
			vertices: 5,
			edges:    [][2]int{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}},
			want:     [][2]int{{0, 3}, {3, 4}},
		},
		{
			name:     "Graph with no bridges",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     [][2]int{},
		},
		{
			name:     "Graph with multiple bridges",
			vertices: 7,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {1, 4}, {1, 6}, {3, 5}, {4, 5}},
			want:     [][2]int{{1, 6}, {1, 3}, {3, 5}},
		},
		{
			name:     "Disconnected graph",
			vertices: 7,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 6}, {6, 3}},
			want:     [][2]int{},
		},
		{
			name:     "Linear graph",
			vertices: 5,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:     [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindBridges()
			sortBridges(got)
			sortBridges(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBridges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortBridges(bridges [][2]int) {
	for i := range bridges {
		if bridges[i][0] > bridges[i][1] {
			bridges[i][0], bridges[i][1] = bridges[i][1], bridges[i][0]
		}
	}
	sort.Slice(bridges, func(i, j int) bool {
		if bridges[i][0] == bridges[j][0] {
			return bridges[i][1] < bridges[j][1]
		}
		return bridges[i][0] < bridges[j][0]
	})
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
