package scc

// Graph represents a directed graph using an adjacency list
type Graph struct {
	V     int
	Edges [][]int
}

// NewGraph creates a new Graph with V vertices
func NewGraph(V int) *Graph {
	return &Graph{
		V:     V,
		Edges: make([][]int, V),
	}
}

// AddEdge adds a directed edge from u to v
func (g *Graph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
}

// FindSCCs finds the strongly connected components in the graph
func (g *Graph) FindSCCs() [][]int {
	stack := make([]int, 0, g.V)
	visited := make([]bool, g.V)

	// First DFS to fill the stack
	for v := 0; v < g.V; v++ {
		if !visited[v] {
			g.fillOrder(v, visited, &stack)
		}
	}

	// Create the transposed graph
	gt := g.getTranspose()

	// Reset visited array
	for i := range visited {
		visited[i] = false
	}

	// Second DFS to find SCCs
	var sccs [][]int
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[v] {
			var component []int
			gt.dfs(v, visited, &component)
			sccs = append(sccs, component)
		}
	}

	return sccs
}

func (g *Graph) fillOrder(v int, visited []bool, stack *[]int) {
	visited[v] = true
	for _, w := range g.Edges[v] {
		if !visited[w] {
			g.fillOrder(w, visited, stack)
		}
	}
	*stack = append(*stack, v)
}

func (g *Graph) getTranspose() *Graph {
	gt := NewGraph(g.V)
	for v := 0; v < g.V; v++ {
		for _, w := range g.Edges[v] {
			gt.AddEdge(w, v)
		}
	}
	return gt
}

func (g *Graph) dfs(v int, visited []bool, component *[]int) {
	visited[v] = true
	*component = append(*component, v)
	for _, w := range g.Edges[v] {
		if !visited[w] {
			g.dfs(w, visited, component)
		}
	}
}
