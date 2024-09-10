package dijkstra

import (
	"container/heap"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents a weighted graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  [][]Edge
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make([][]Edge, vertices),
	}
}

// AddEdge adds a weighted edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
}

// Dijkstra runs Dijkstra's algorithm to find the shortest paths from a source vertex
func (g *Graph) Dijkstra(source int) ([]int, []int) {
	dist := make([]int, g.Vertices)
	prev := make([]int, g.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Value: source, Priority: 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).Value

		for _, edge := range g.AdjList[u] {
			v := edge.To
			alt := dist[u] + edge.Weight

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(&pq, &Item{Value: v, Priority: alt})
			}
		}
	}

	return dist, prev
}

// GetPath retrieves the shortest path from the source to a destination
func (g *Graph) GetPath(source, destination int) []int {
	_, prev := g.Dijkstra(source)
	path := []int{}
	for v := destination; v != -1; v = prev[v] {
		path = append([]int{v}, path...)
	}
	if path[0] != source {
		return nil // No path found
	}
	return path
}

// Item is an item in the priority queue
type Item struct {
	Value    int
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
