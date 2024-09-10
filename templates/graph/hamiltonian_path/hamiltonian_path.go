package hamiltonianpath

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

// FindHamiltonianPath finds a Hamiltonian path in the graph if one exists
func (g *Graph) FindHamiltonianPath() []int {
	path := make([]int, g.V)
	for i := range path {
		path[i] = -1
	}

	// Try starting from each vertex
	for start := 0; start < g.V; start++ {
		path[0] = start
		if g.hamiltonianPathUtil(path, 1) {
			return path
		}
		path[0] = -1
	}

	return nil
}

func (g *Graph) hamiltonianPathUtil(path []int, pos int) bool {
	if pos == g.V {
		return true
	}

	for v := 0; v < g.V; v++ {
		if g.isSafe(v, path, pos) {
			path[pos] = v

			if g.hamiltonianPathUtil(path, pos+1) {
				return true
			}

			path[pos] = -1
		}
	}

	return false
}

func (g *Graph) isSafe(v int, path []int, pos int) bool {
	// Check if this vertex is adjacent to the previously added vertex
	if pos > 0 && !contains(g.Adj[path[pos-1]], v) {
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
