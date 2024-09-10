package bellmanford

import (
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	Source      int
	Destination int
	Weight      int
}

// Graph represents a weighted graph
type Graph struct {
	Vertices int
	Edges    []Edge
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    []Edge{},
	}
}

// AddEdge adds a weighted edge to the graph
func (g *Graph) AddEdge(source, destination, weight int) {
	g.Edges = append(g.Edges, Edge{Source: source, Destination: destination, Weight: weight})
}

// BellmanFord runs the Bellman-Ford algorithm to find the shortest paths from a source vertex
// It returns the distances to all vertices and a boolean indicating if there's a negative cycle
func (g *Graph) BellmanFord(source int) ([]int, bool) {
	dist := make([]int, g.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	// Relax all edges |V| - 1 times
	for i := 1; i < g.Vertices; i++ {
		for _, edge := range g.Edges {
			if dist[edge.Source] != math.MaxInt32 && dist[edge.Source]+edge.Weight < dist[edge.Destination] {
				dist[edge.Destination] = dist[edge.Source] + edge.Weight
			}
		}
	}

	// Check for negative-weight cycles
	for _, edge := range g.Edges {
		if dist[edge.Source] != math.MaxInt32 && dist[edge.Source]+edge.Weight < dist[edge.Destination] {
			return nil, true // Negative cycle detected
		}
	}

	return dist, false
}

// GetPath retrieves the shortest path from the source to a destination
func (g *Graph) GetPath(source, destination int) []int {
	dist, hasNegativeCycle := g.BellmanFord(source)
	if hasNegativeCycle {
		return nil // No valid path due to negative cycle
	}

	path := []int{destination}
	for current := destination; current != source; {
		found := false
		for _, edge := range g.Edges {
			if edge.Destination == current && dist[edge.Source]+edge.Weight == dist[current] {
				path = append([]int{edge.Source}, path...)
				current = edge.Source
				found = true
				break
			}
		}
		if !found {
			return nil // No path found
		}
	}

	return path
}
