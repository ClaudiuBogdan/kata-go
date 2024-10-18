package kruskalmst

import "sort"

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
	return &Graph{
		V:     V,
		Edges: make([]Edge, 0),
	}
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) error {
	g.Edges = append(g.Edges, Edge{Src: src, Dest: dest, Weight: weight})
	return nil
}

// UnionFind represents a disjoint set data structure
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new UnionFind data structure
func NewUnionFind(size int) *UnionFind {

	parent := make([]int, size)

	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	return &UnionFind{
		rank:   make([]int, size),
		parent: parent,
	}
}

// Find finds the parent of a vertex with path compression
func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}

	return uf.parent[i]
}

// Union performs union of two sets of vertices by rank
func (uf *UnionFind) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)

	// TODO: get the highest rank
	if uf.rank[xRoot] > uf.rank[yRoot] {
		uf.parent[yRoot] = xRoot
	} else if uf.rank[xRoot] < uf.rank[yRoot] {
		uf.parent[xRoot] = yRoot
	} else {
		uf.rank[xRoot]++
		uf.parent[yRoot] = xRoot
	}
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	// Sort the edges
	// Connect the edges that are not on same set
	mst := make([]Edge, 0)
	uf := NewUnionFind(g.V)

	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	for _, edge := range g.Edges {
		srcRoot := uf.Find(edge.Src)
		destRoot := uf.Find(edge.Dest)

		if srcRoot != destRoot {
			mst = append(mst, edge)
			uf.Union(srcRoot, destRoot)
		}
	}

	return mst
}
