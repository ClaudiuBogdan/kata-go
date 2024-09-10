package astar

import (
	"reflect"
	"testing"
)

func TestAStar(t *testing.T) {
	tests := []struct {
		name      string
		gridSetup func(*Grid)
		start     *Node
		goal      *Node
		want      [][2]int
	}{
		{
			name: "Simple path",
			gridSetup: func(g *Grid) {},
			start:     &Node{X: 0, Y: 0},
			goal:      &Node{X: 2, Y: 2},
			want:      [][2]int{{0, 0}, {1, 1}, {2, 2}},
		},
		{
			name: "Path with obstacle",
			gridSetup: func(g *Grid) {
				g.Nodes[1][1].Walkable = false
			},
			start: &Node{X: 0, Y: 0},
			goal:  &Node{X: 2, Y: 2},
			want:  [][2]int{{0, 0}, {1, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "No path",
			gridSetup: func(g *Grid) {
				g.Nodes[1][0].Walkable = false
				g.Nodes[1][1].Walkable = false
				g.Nodes[1][2].Walkable = false
			},
			start: &Node{X: 0, Y: 0},
			goal:  &Node{X: 2, Y: 2},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid(3, 3)
			tt.gridSetup(grid)

			start := grid.Nodes[tt.start.Y][tt.start.X]
			goal := grid.Nodes[tt.goal.Y][tt.goal.X]

			path := AStar(grid, start, goal)

			if tt.want == nil && path != nil {
				t.Errorf("Expected no path, but got a path")
				return
			}

			if tt.want != nil && path == nil {
				t.Errorf("Expected a path, but got no path")
				return
			}

			if tt.want != nil {
				got := make([][2]int, len(path))
				for i, node := range path {
					got[i] = [2]int{node.X, node.Y}
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AStar() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestHeuristic(t *testing.T) {
	tests := []struct {
		name string
		node *Node
		goal *Node
		want float64
	}{
		{
			name: "Same point",
			node: &Node{X: 0, Y: 0},
			goal: &Node{X: 0, Y: 0},
			want: 0,
		},
		{
			name: "Horizontal",
			node: &Node{X: 0, Y: 0},
			goal: &Node{X: 3, Y: 0},
			want: 3,
		},
		{
			name: "Vertical",
			node: &Node{X: 0, Y: 0},
			goal: &Node{X: 0, Y: 4},
			want: 4,
		},
		{
			name: "Diagonal",
			node: &Node{X: 0, Y: 0},
			goal: &Node{X: 3, Y: 4},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Heuristic(tt.node, tt.goal)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("Heuristic() = %v, want %v", got, tt.want)
			}
		})
	}
}
