package kruskalmst

import (
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
func (g *Graph) AddEdge(src, dest, weight int) {
	g.Edges = append(g.Edges, Edge{src, dest, weight})
}

// KruskalMST finds the Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) KruskalMST() []Edge {
	// Sort edges by weight
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	// Initialize disjoint set
	parent := make([]int, g.V)
	rank := make([]int, g.V)
	for i := range parent {
		parent[i] = i
		rank[i] = 0
	}

	mst := []Edge{}

	for _, edge := range g.Edges {
		x := find(parent, edge.Src)
		y := find(parent, edge.Dest)

		if x != y {
			mst = append(mst, edge)
			union(parent, rank, x, y)
		}

		if len(mst) == g.V-1 {
			break
		}
	}

	return mst
}

// find finds the parent of a vertex with path compression
func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

// union performs union of two sets of vertices by rank
func union(parent, rank []int, x, y int) {
	xroot := find(parent, x)
	yroot := find(parent, y)

	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}
