package toposort

// Graph represents a directed acyclic graph using an adjacency list
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

// TopologicalSort performs topological sorting on the graph
func (g *Graph) TopologicalSort() ([]int, bool) {
	visited := make([]bool, g.V)
	stack := make([]int, 0, g.V)
	recStack := make([]bool, g.V)

	for v := 0; v < g.V; v++ {
		if !visited[v] {
			if g.dfs(v, visited, &stack, recStack) {
				// Cycle detected
				return nil, false
			}
		}
	}

	// Reverse the stack to get the topological order
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack, true
}

func (g *Graph) dfs(v int, visited []bool, stack *[]int, recStack []bool) bool {
	visited[v] = true
	recStack[v] = true

	for _, w := range g.Edges[v] {
		if !visited[w] {
			if g.dfs(w, visited, stack, recStack) {
				return true
			}
		} else if recStack[w] {
			// Cycle detected
			return true
		}
	}

	recStack[v] = false
	*stack = append(*stack, v)
	return false
}
