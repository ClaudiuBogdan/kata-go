package toposort

// Graph represents a directed graph using an adjacency list.
type Graph struct {
	V     int     // Number of vertices
	Edges [][]int // Adjacency list
}

// NewGraph creates a new graph with V vertices.
func NewGraph(V int) *Graph {

}

// AddEdge adds a directed edge from v to w.
func (g *Graph) AddEdge(v, w int) {

}

// TopologicalSort performs a topological sort of the graph.
// It returns the sorted order and an error if a cycle is detected.
func (g *Graph) TopologicalSort() ([]int, error) {

}
