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
	return &Graph{V, []Edge{}}
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) error {
	if src < 0 || src >= g.V || dest < 0 || dest >= g.V {
		return errors.New("invalid args")
	}

	if weight < 0 {
		return errors.New("invalid weight")
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
	uf := &UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}

	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
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

	if uf.rank[xroot] > uf.rank[yroot] {
		uf.parent[yroot] = xroot
	} else if uf.rank[xroot] < uf.rank[yroot] {
		uf.parent[xroot] = yroot
	} else {
		uf.parent[yroot] = xroot
		uf.rank[xroot]++
	}
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	uf := NewUnionFind(g.V)
	mst := make([]Edge, 0)

	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	for _, edge := range g.Edges {
		xroot := uf.Find(edge.Src)
		yroot := uf.Find(edge.Dest)

		if xroot != yroot {
			mst = append(mst, edge)
			uf.Union(xroot, yroot)
		}

		if len(mst) == g.V-1 {
			break
		}
	}

	return mst
}
