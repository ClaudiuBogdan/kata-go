package mincut

import (
	"math"
)

// Graph represents an undirected graph using an adjacency matrix
type Graph struct {
	V     int     // Number of vertices
	Edges [][]int // Adjacency matrix
}

// NewGraph creates a new Graph with V vertices
func NewGraph(V int) *Graph {
	edges := make([][]int, V)
	for i := range edges {
		edges[i] = make([]int, V)
	}
	return &Graph{
		V:     V,
		Edges: edges,
	}
}

// AddEdge adds an undirected edge between vertices u and v with weight w
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u][v] += w
	g.Edges[v][u] += w
}

// FindMinCut finds the minimum cut in the graph using the Stoer-Wagner algorithm
func (g *Graph) FindMinCut() (int, []int) {
	n := g.V
	co := make([][]int, n)
	for i := range co {
		co[i] = []int{i}
	}

	bestCut := math.MaxInt32
	var bestPartition []int

	for n > 1 {
		cut, s, t := g.minimumCutPhase(n)
		if cut < bestCut {
			bestCut = cut
			bestPartition = make([]int, len(co[t]))
			copy(bestPartition, co[t])
		}

		// Merge vertices s and t
		co[s] = append(co[s], co[t]...)
		for i := 0; i < n; i++ {
			g.Edges[s][i] += g.Edges[t][i]
			g.Edges[i][s] = g.Edges[s][i]
		}

		// Remove vertex t
		g.Edges[t] = g.Edges[n-1]
		for i := 0; i < n; i++ {
			g.Edges[i][t] = g.Edges[i][n-1]
		}
		co[t] = co[n-1]
		n--
	}

	return bestCut, bestPartition
}

func (g *Graph) minimumCutPhase(n int) (int, int, int) {
	added := make([]bool, n)
	w := make([]int, n)
	var a, s, t int

	for i := 0; i < n-1; i++ {
		added[a] = true
		s = t
		t = -1
		for j := 0; j < n; j++ {
			if !added[j] {
				w[j] += g.Edges[a][j]
				if t == -1 || w[j] > w[t] {
					t = j
				}
			}
		}
		a = t
	}

	return w[t], s, t
}
