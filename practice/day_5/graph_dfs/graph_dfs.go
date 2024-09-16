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
	result := make([]int, 0, g.Vertices)
	stack := make([]int, 0, g.Vertices)

	// Add first element to stack
	stack = append(stack, source)

	for len(stack) > 0 {
		// Pop from the stack
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[v] {
			visited[v] = true
			result = append(result, v)

			// Iterate over neighbors
			neighbors := g.AdjList[v]
			n := len(neighbors)
			// TODO: try to visualize this in normal and reverse order
			for i := n - 1; i >= 0; i-- {
				neighbor := neighbors[i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	return result
}

// HasPath checks if there's a path between source and destination using DFS
func (g *Graph) HasPath(source, destination int) bool {
	visited := make(map[int]bool)
	stack := make([]int, 0, g.Vertices)

	stack = append(stack, source)

	for len(stack) > 0 {
		// Pop from the stack
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
	allPaths := [][]int{}
	visited := make(map[int]bool)
	path := []int{source}

	var dfs func(int) // declare dfs here to call it recursively from withing dfs
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

	// Call dfs here
	dfs(source)

	return allPaths
}
