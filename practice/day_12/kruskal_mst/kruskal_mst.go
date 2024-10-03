package kruskalmst

import (
	"fmt"
	"sort"
)

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
		Edges: make([]Edge, 0, V),
	}
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) error {
	if src < 0 || dest >= g.V {
		return fmt.Errorf("invalid vertices")
	}

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
	rank := make([]int, size)

	for i := range parent {
		parent[i] = i
	}

	return &UnionFind{
		parent,
		rank,
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
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[x] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[y] = rootX
	} else {
		uf.rank[rootX]++
		// TODO assign parent of rootY to rootX, not parent of y to rootX
		uf.parent[rootY] = rootX
	}
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	// Sort the vertices, smaller first
	// Check if the vertex is not part of the set. Build the tree
	uf := NewUnionFind(g.V)
	mst := make([]Edge, 0)

	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	for _, edge := range g.Edges {
		if uf.Find(edge.Src) != uf.Find(edge.Dest) {
			mst = append(mst, edge)
			uf.Union(edge.Src, edge.Dest)
		}
	}

	return mst
}
