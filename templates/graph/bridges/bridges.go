package bridges

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.AdjList[v] = append(g.AdjList[v], w)
	g.AdjList[w] = append(g.AdjList[w], v)
}

// FindBridges finds all bridges in the graph
func (g *Graph) FindBridges() [][2]int {
	visited := make([]bool, g.Vertices)
	disc := make([]int, g.Vertices)
	low := make([]int, g.Vertices)
	parent := make([]int, g.Vertices)
	bridges := [][2]int{}

	for i := range parent {
		parent[i] = -1
	}

	var time int

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		disc[u] = time
		low[u] = time
		time++

		for _, v := range g.AdjList[u] {
			if !visited[v] {
				parent[v] = u
				dfs(v)

				low[u] = min(low[u], low[v])

				if low[v] > disc[u] {
					bridges = append(bridges, [2]int{u, v})
				}
			} else if v != parent[u] {
				low[u] = min(low[u], disc[v])
			}
		}
	}

	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	return bridges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
