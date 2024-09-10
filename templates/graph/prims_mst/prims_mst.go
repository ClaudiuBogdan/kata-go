package primsmst

import (
	"container/heap"
)

// Edge represents an edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	V     int
	Edges [][]Edge
}

// NewGraph creates a new Graph with V vertices
func NewGraph(V int) *Graph {
	return &Graph{
		V:     V,
		Edges: make([][]Edge, V),
	}
}

// AddEdge adds an undirected edge between vertices u and v with weight w
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{To: v, Weight: w})
	g.Edges[v] = append(g.Edges[v], Edge{To: u, Weight: w})
}

// PrimMST finds the Minimum Spanning Tree using Prim's algorithm
func (g *Graph) PrimMST() []Edge {
	mst := []Edge{}
	visited := make([]bool, g.V)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Start from vertex 0
	visited[0] = true
	for _, e := range g.Edges[0] {
		heap.Push(&pq, &Item{Edge: e, From: 0, Priority: e.Weight})
	}

	for len(mst) < g.V-1 && pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if visited[item.Edge.To] {
			continue
		}

		visited[item.Edge.To] = true
		mst = append(mst, Edge{To: item.Edge.To, Weight: item.Edge.Weight})

		for _, e := range g.Edges[item.Edge.To] {
			if !visited[e.To] {
				heap.Push(&pq, &Item{Edge: e, From: item.Edge.To, Priority: e.Weight})
			}
		}
	}

	return mst
}

// Item is an item in the priority queue
type Item struct {
	Edge     Edge
	From     int
	Priority int
	Index    int
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
