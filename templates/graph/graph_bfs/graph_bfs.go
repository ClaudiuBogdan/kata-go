package bfs

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {

}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {

}

// BFS performs a breadth-first search starting from the given source vertex
func (g *Graph) BFS(source int) []int {

}

// ShortestPath finds the shortest path between source and destination using BFS
func (g *Graph) ShortestPath(source, destination int) []int {

}

// reconstructPath reconstructs the path from source to destination using the parent array
func (g *Graph) reconstructPath(parent []int, source, destination int) []int {

}
