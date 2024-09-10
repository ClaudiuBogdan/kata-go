package hamiltoniancycle

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	V   int
	Adj [][]int
}

// NewGraph creates a new Graph with V vertices
func NewGraph(V int) *Graph {
	return &Graph{
		V:   V,
		Adj: make([][]int, V),
	}
}

// AddEdge adds an undirected edge between vertices v and w
func (g *Graph) AddEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w)
	g.Adj[w] = append(g.Adj[w], v)
}

// FindHamiltonianCycle finds a Hamiltonian cycle in the graph if one exists
func (g *Graph) FindHamiltonianCycle() []int {
	path := make([]int, g.V)
	for i := range path {
		path[i] = -1
	}
	path[0] = 0

	if g.hamiltonianCycleUtil(path, 1) {
		return path
	}
	return nil
}

func (g *Graph) hamiltonianCycleUtil(path []int, pos int) bool {
	if pos == g.V {
		// Check if there is an edge from the last vertex to the first vertex
		for _, v := range g.Adj[path[pos-1]] {
			if v == path[0] {
				return true
			}
		}
		return false
	}

	for v := 1; v < g.V; v++ {
		if g.isSafe(v, path, pos) {
			path[pos] = v

			if g.hamiltonianCycleUtil(path, pos+1) {
				return true
			}

			path[pos] = -1
		}
	}

	return false
}

func (g *Graph) isSafe(v int, path []int, pos int) bool {
	// Check if this vertex is adjacent to the previously added vertex
	if !contains(g.Adj[path[pos-1]], v) {
		return false
	}

	// Check if the vertex has already been included in the path
	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}

	return true
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
