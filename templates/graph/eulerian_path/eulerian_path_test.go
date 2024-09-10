package eulerianpath

import (
	"testing"
)

func TestHasEulerianPath(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		want     bool
	}{
		{
			name:     "Eulerian path (not circuit)",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}},
			want:     true,
		},
		{
			name:     "Eulerian circuit",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			want:     true,
		},
		{
			name:     "No Eulerian path",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			if got := g.HasEulerianPath(); got != tt.want {
				t.Errorf("HasEulerianPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindEulerianPath(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][2]int
		wantPath bool
	}{
		{
			name:     "Eulerian path (not circuit)",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}},
			wantPath: true,
		},
		{
			name:     "Eulerian circuit",
			vertices: 3,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 0}},
			wantPath: true,
		},
		{
			name:     "No Eulerian path",
			vertices: 4,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}},
			wantPath: false,
		},
		{
			name:     "Empty graph",
			vertices: 3,
			edges:    [][2]int{},
			wantPath: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.vertices)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1])
			}

			got := g.FindEulerianPath()

			if (got != nil) != tt.wantPath {
				t.Errorf("FindEulerianPath() returned path = %v, want path = %v", got != nil, tt.wantPath)
			}

			if got != nil && !isValidEulerianPath(g, got) {
				t.Errorf("FindEulerianPath() returned an invalid Eulerian path: %v", got)
			}
		})
	}
}

// isValidEulerianPath checks if the given path is a valid Eulerian path
func isValidEulerianPath(g *Graph, path []int) bool {
	if len(path) == 0 {
		return len(g.AdjList) == 0
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

	for i := 0; i < len(path)-1; i++ {
		v, u := path[i], path[i+1]
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
