package maxflow

import (
	"container/list"
	"math"
)

// FlowNetwork represents a flow network
type FlowNetwork struct {
	V     int       // Number of vertices
	Graph [][]int   // Residual graph
	Flow  [][]int   // Flow graph
}

// NewFlowNetwork creates a new FlowNetwork with V vertices
func NewFlowNetwork(V int) *FlowNetwork {
	graph := make([][]int, V)
	flow := make([][]int, V)
	for i := range graph {
		graph[i] = make([]int, V)
		flow[i] = make([]int, V)
	}
	return &FlowNetwork{
		V:     V,
		Graph: graph,
		Flow:  flow,
	}
}

// AddEdge adds a directed edge from u to v with capacity c
func (fn *FlowNetwork) AddEdge(u, v, c int) {
	fn.Graph[u][v] = c
}

// FindMaxFlow finds the maximum flow from source to sink
func (fn *FlowNetwork) FindMaxFlow(source, sink int) int {
	maxFlow := 0

	for {
		// Find an augmenting path
		parent := fn.bfs(source, sink)
		if parent == nil {
			break // No augmenting path found, we're done
		}

		// Find the minimum residual capacity along the path
		pathFlow := math.MaxInt32
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			pathFlow = min(pathFlow, fn.Graph[u][v]-fn.Flow[u][v])
		}

		// Update residual capacities and reverse edges along the path
		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			fn.Flow[u][v] += pathFlow
			fn.Flow[v][u] -= pathFlow
		}

		maxFlow += pathFlow
	}

	return maxFlow
}

// bfs performs breadth-first search to find an augmenting path
func (fn *FlowNetwork) bfs(source, sink int) []int {
	parent := make([]int, fn.V)
	for i := range parent {
		parent[i] = -1
	}
	parent[source] = source

	queue := list.New()
	queue.PushBack(source)

	for queue.Len() > 0 {
		u := queue.Remove(queue.Front()).(int)
		for v := 0; v < fn.V; v++ {
			if parent[v] == -1 && fn.Graph[u][v] > fn.Flow[u][v] {
				parent[v] = u
				if v == sink {
					return parent
				}
				queue.PushBack(v)
			}
		}
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
