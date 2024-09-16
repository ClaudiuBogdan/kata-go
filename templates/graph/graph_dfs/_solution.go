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
	var result []int
	stack := []int{source}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[v] {
			visited[v] = true
			result = append(result, v)

			for i := len(g.AdjList[v]) - 1; i >= 0; i-- {
				if !visited[g.AdjList[v][i]] {
					stack = append(stack, g.AdjList[v][i])
				}
			}
		}
	}

	return result
}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {
	visited := make(map[int]bool)
	stack := []int{source}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if v == destination {
			return true
		}

		if !visited[v] {
			visited[v] = true
			for _, neighbor := range g.AdjList[v] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	return false
}

// FindAllPaths finds all paths between source and destination using DFS
func (g *Graph) FindAllPaths(source, destination int) [][]int {
	var allPaths [][]int
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
				path = path[:len(path)-1]
			}
		}
		visited[v] = false
	}

	dfs(source)
	return allPaths
}