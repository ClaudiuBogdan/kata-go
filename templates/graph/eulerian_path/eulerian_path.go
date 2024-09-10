package eulerianpath

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

// HasEulerianPath checks if the graph has an Eulerian path
func (g *Graph) HasEulerianPath() bool {
	if !g.isConnected() {
		return false
	}

	oddDegreeCount := 0
	for _, edges := range g.AdjList {
		if len(edges)%2 != 0 {
			oddDegreeCount++
		}
	}

	return oddDegreeCount == 0 || oddDegreeCount == 2
}

// FindEulerianPath finds an Eulerian path in the graph if one exists
func (g *Graph) FindEulerianPath() []int {
	if !g.HasEulerianPath() {
		return nil
	}

	// Find a starting vertex (odd degree vertex if exists, otherwise any vertex)
	start := g.findStartVertex()

	// Create a copy of the graph
	tempGraph := g.clone()

	path := []int{}
	var dfs func(v int)
	dfs = func(v int) {
		for len(tempGraph.AdjList[v]) > 0 {
			u := tempGraph.AdjList[v][0]
			tempGraph.removeEdge(v, u)
			dfs(u)
		}
		path = append([]int{v}, path...)
	}

	dfs(start)

	return path
}

// isConnected checks if the graph is connected
func (g *Graph) isConnected() bool {
	visited := make(map[int]bool)
	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		for _, u := range g.AdjList[v] {
			if !visited[u] {
				dfs(u)
			}
		}
	}

	// Start DFS from any vertex
	start := g.findStartVertex()
	dfs(start)

	return len(visited) == len(g.AdjList)
}

// findStartVertex finds a suitable starting vertex for the Eulerian path
func (g *Graph) findStartVertex() int {
	start := 0
	for v, edges := range g.AdjList {
		if len(edges)%2 != 0 {
			return v
		}
		start = v
	}
	return start
}

// clone creates a deep copy of the graph
func (g *Graph) clone() *Graph {
	newGraph := NewGraph(g.Vertices)
	for v, edges := range g.AdjList {
		newGraph.AdjList[v] = append([]int{}, edges...)
	}
	return newGraph
}

// removeEdge removes an edge from the graph
func (g *Graph) removeEdge(v, u int) {
	g.AdjList[v] = removeFromSlice(g.AdjList[v], u)
	g.AdjList[u] = removeFromSlice(g.AdjList[u], v)
}

// removeFromSlice removes the first occurrence of a value from a slice
func removeFromSlice(s []int, val int) []int {
	for i, v := range s {
		if v == val {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
