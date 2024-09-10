package floydwarshall

import (
	"math"
)

// Graph represents a weighted graph using an adjacency matrix
type Graph struct {
	Vertices int
	Matrix   [][]float64
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	matrix := make([][]float64, vertices)
	for i := range matrix {
		matrix[i] = make([]float64, vertices)
		for j := range matrix[i] {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = math.Inf(1)
			}
		}
	}
	return &Graph{
		Vertices: vertices,
		Matrix:   matrix,
	}
}

// AddEdge adds a weighted edge to the graph
func (g *Graph) AddEdge(from, to int, weight float64) {
	g.Matrix[from][to] = weight
}

// FloydWarshall runs the Floyd-Warshall algorithm to find all-pairs shortest paths
func (g *Graph) FloydWarshall() [][]float64 {
	dist := make([][]float64, g.Vertices)
	for i := range dist {
		dist[i] = make([]float64, g.Vertices)
		copy(dist[i], g.Matrix[i])
	}

	for k := 0; k < g.Vertices; k++ {
		for i := 0; i < g.Vertices; i++ {
			for j := 0; j < g.Vertices; j++ {
				if dist[i][k] != math.Inf(1) && dist[k][j] != math.Inf(1) {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
					}
				}
			}
		}
	}

	return dist
}

// GetPath retrieves the shortest path between two vertices
func (g *Graph) GetPath(from, to int) []int {
	dist := g.FloydWarshall()
	if dist[from][to] == math.Inf(1) {
		return nil // No path exists
	}

	path := []int{from}
	for from != to {
		for next := 0; next < g.Vertices; next++ {
			if next != from && g.Matrix[from][next] != math.Inf(1) &&
				dist[from][to] == g.Matrix[from][next]+dist[next][to] {
				path = append(path, next)
				from = next
				break
			}
		}
	}

	return path
}

// HasNegativeCycle checks if the graph contains a negative cycle
func (g *Graph) HasNegativeCycle() bool {
	dist := g.FloydWarshall()
	for i := 0; i < g.Vertices; i++ {
		if dist[i][i] < 0 {
			return true
		}
	}
	return false
}
