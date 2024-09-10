package bipartitematching

// BipartiteGraph represents a bipartite graph
type BipartiteGraph struct {
	U, V  int     // Number of vertices in sets U and V
	Edges [][]int // Adjacency list representation
}

// NewBipartiteGraph creates a new BipartiteGraph with u vertices in set U and v vertices in set V
func NewBipartiteGraph(u, v int) *BipartiteGraph {
	return &BipartiteGraph{
		U:     u,
		V:     v,
		Edges: make([][]int, u),
	}
}

// AddEdge adds an edge from vertex u in set U to vertex v in set V
func (g *BipartiteGraph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
}

// FindMaximumMatching finds the maximum matching in the bipartite graph
func (g *BipartiteGraph) FindMaximumMatching() []int {
	matching := make([]int, g.V)
	for i := range matching {
		matching[i] = -1
	}

	result := 0
	for u := 0; u < g.U; u++ {
		visited := make([]bool, g.V)
		if g.dfs(u, visited, matching) {
			result++
		}
	}

	return matching
}

// dfs performs depth-first search to find an augmenting path
func (g *BipartiteGraph) dfs(u int, visited []bool, matching []int) bool {
	for _, v := range g.Edges[u] {
		if !visited[v] {
			visited[v] = true

			if matching[v] == -1 || g.dfs(matching[v], visited, matching) {
				matching[v] = u
				return true
			}
		}
	}
	return false
}

// GetMatchingSize returns the size of the maximum matching
func (g *BipartiteGraph) GetMatchingSize(matching []int) int {
	size := 0
	for _, m := range matching {
		if m != -1 {
			size++
		}
	}
	return size
}
