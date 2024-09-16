package kruskalmst

import (
	"errors"
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
		Edges: []Edge{},
	}
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) error {
	if src < 0 || src >= g.V || dest < 0 || dest >= g.V {
		return errors.New("invalid vertex index")
	}
	if weight < 0 {
		return errors.New("negative weight not allowed")
	}
	g.Edges = append(g.Edges, Edge{src, dest, weight})
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
	return &UnionFind{parent, rank}
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
	xroot := uf.Find(x)
	yroot := uf.Find(y)

	if uf.rank[xroot] < uf.rank[yroot] {
		uf.parent[xroot] = yroot
	} else if uf.rank[xroot] > uf.rank[yroot] {
		uf.parent[yroot] = xroot
	} else {
		uf.parent[yroot] = xroot
		uf.rank[xroot]++
	}
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	// Sort edges by weight
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	uf := NewUnionFind(g.V)
	mst := []Edge{}

	for _, edge := range g.Edges {
		x := uf.Find(edge.Src)
		y := uf.Find(edge.Dest)

		if x != y {
			mst = append(mst, edge)
			uf.Union(x, y)
		}

		if len(mst) == g.V-1 {
			break
		}
	}


	return mst
}