package articulationpoints

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

// FindArticulationPoints finds all articulation points in the graph
func (g *Graph) FindArticulationPoints() []int {
	visited := make([]bool, g.Vertices)
	disc := make([]int, g.Vertices)
	low := make([]int, g.Vertices)
	parent := make([]int, g.Vertices)
	articulationPoints := make([]bool, g.Vertices)

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
		childCount := 0

		for _, v := range g.AdjList[u] {
			if !visited[v] {
				childCount++
				parent[v] = u
				dfs(v)

				low[u] = min(low[u], low[v])

				if parent[u] == -1 && childCount > 1 {
					articulationPoints[u] = true
				}

				if parent[u] != -1 && low[v] >= disc[u] {
					articulationPoints[u] = true
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

	result := []int{}
	for i, isArticulation := range articulationPoints {
		if isArticulation {
			result = append(result, i)
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
