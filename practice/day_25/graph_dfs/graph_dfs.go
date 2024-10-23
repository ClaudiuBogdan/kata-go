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
	visited := make(map[int]bool)
	result := make([]int, 0)

	var dfs func(current int)

	dfs = func(current int) {
		result = append(result, current)

		for _, neighbor := range g.AdjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				dfs(neighbor)
			}
		}
	}

	visited[source] = true
	dfs(source)

	return result
}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {
	visited := make(map[int]bool)

	var dfs func(current int) bool

	dfs = func(current int) bool {
		if current == destination {
			return true
		}

		for _, neighbor := range g.AdjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				found := dfs(neighbor)
				if found {
					return true
				}
			}
		}
		return false
	}

	visited[source] = true
	return dfs(source)
}

// FindAllPaths finds all paths between source and destination using DFS
func (g *Graph) FindAllPaths(source, destination int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	visited := make(map[int]bool)

	var dfs func(current int)

	dfs = func(current int) {
		path = append(path, current)
		visited[current] = true

		// TODO: use if else
		if current == destination {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		} else {
			for _, neighbor := range g.AdjList[current] {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}

		path = path[:len(path)-1]
		visited[current] = false
	}

	dfs(source)

	return result
}
