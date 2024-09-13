package dfs

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// New creates a new Graph with the given number of vertices
func New(vertices int) *Graph {
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

// DFS performs a depth-first search starting from the given source vertex
func (g *Graph) DFS(source int) []int {
	visited := make([]bool, g.Vertices)
	result := []int{}

	var dfsUtil func(v int)
	dfsUtil = func(v int) {
		visited[v] = true
		result = append(result, v)

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				dfsUtil(neighbor)
			}
		}
	}

	dfsUtil(source)
	return result
}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {
	visited := make([]bool, g.Vertices)

	var dfsUtil func(v int) bool
	dfsUtil = func(v int) bool {
		if v == destination {
			return true
		}

		visited[v] = true

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				if dfsUtil(neighbor) {
					return true
				}
			}
		}

		return false
	}

	return dfsUtil(source)
}

// FindAllPaths finds all paths between source and destination using DFS
func (g *Graph) FindAllPaths(source, destination int) [][]int {
	visited := make([]bool, g.Vertices)
	currentPath := []int{}
	allPaths := [][]int{}

	var dfsUtil func(v int)
	dfsUtil = func(v int) {
		visited[v] = true
		currentPath = append(currentPath, v)

		if v == destination {
			pathCopy := make([]int, len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, pathCopy)
		} else {
			for _, neighbor := range g.AdjList[v] {
				if !visited[neighbor] {
					dfsUtil(neighbor)
				}
			}
		}

		visited[v] = false
		currentPath = currentPath[:len(currentPath)-1]
	}

	dfsUtil(source)
	return allPaths
}
