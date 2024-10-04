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
		AdjList:  make(map[int][]int, vertices),
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.AdjList[v] = append(g.AdjList[v], w)
	g.AdjList[w] = append(g.AdjList[w], v)
}

// DFS performs a depth-first search starting from the given source vertex
func (g *Graph) DFS(source int) []int {
	visited := make(map[int]bool)
	path := make([]int, 0)

	var dfs func(v int)

	visited[source] = true

	dfs = func(v int) {
		visited[v] = true
		path = append(path, v)

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(source)

	return path
}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {
	visited := make(map[int]bool)

	var dfs func(v int) bool

	dfs = func(v int) bool {
		if v == destination {
			return true
		}
		// TODO: backtrack when checking the path
		visited[v] = true
		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}

		visited[v] = false

		return false
	}

	return dfs(source)
}

// FindAllPaths finds all paths between source and destination using DFS
func (g *Graph) FindAllPaths(source, destination int) [][]int {
	// TODO: this algorithm uses backtracing to generate the paths.
	allPaths := make([][]int, 0)
	visited := make(map[int]bool)
	path := []int{source}

	var dfs func(v int)

	dfs = func(v int) {

		if v == destination {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[v] = true

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				path = append(path, neighbor)
				dfs(neighbor)
				// Backtrack
				path = path[:len(path)-1]
			}
		}
		
		// Backtrack
		visited[v] = false
	}

	dfs(source)

	return allPaths
}
