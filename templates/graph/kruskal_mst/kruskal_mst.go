package kruskalmst

// Edge represents an edge in the graph
type Edge struct {
	Src    int
	Dest   int
	Weight int
}

// Graph represents a weighted undirected graph
type Graph struct {
	V     int
	Edges []Edge
}

// NewGraph creates a new Graph with V vertices
func NewGraph(V int) *Graph {
	
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) error {
	
}

// UnionFind represents a disjoint set data structure
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new UnionFind data structure
func NewUnionFind(size int) *UnionFind {

}

// Find finds the parent of a vertex with path compression
func (uf *UnionFind) Find(i int) int {

}

// Union performs union of two sets of vertices by rank
func (uf *UnionFind) Union(x, y int) {

}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	
}