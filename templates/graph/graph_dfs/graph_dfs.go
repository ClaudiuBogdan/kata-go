package dfs

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// New creates a new Graph with the given number of vertices
func New(vertices int) *Graph {

}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {

}

// DFS performs a depth-first search starting from the given source vertex
func (g *Graph) DFS(source int) []int {

}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {

}

// FindAllPaths finds all paths between source and destination using DFS
func (g *Graph) FindAllPaths(source, destination int) [][]int {

}
