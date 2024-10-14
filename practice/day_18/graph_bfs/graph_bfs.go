package bfs

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList: make(map[int][]int),
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.AdjList[v] = append(g.AdjList[v], w)
	g.AdjList[w] = append(g.AdjList[w], v)
}

// BFS performs a breadth-first search starting from the given source vertex
func (g *Graph) BFS(source int) []int {
	result := make([]int, 0)
	visited := make(map[int]bool)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current)
		
		for _, neighbor := range g.AdjList[current]{
			if !visited[neighbor]{
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return result
}

// ShortestPath finds the shortest path between source and destination using BFS
func (g *Graph) ShortestPath(source, destination int) []int {
	parent := make([]int, g.Vertices)
	visited := make(map[int]bool)

	for i := range parent {
		parent[i] = -1
	}

	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		// TODO: remember to dequeue
		queue = queue[1:]

		if current == destination {
			return g.reconstructPath(parent, destination)
		}

		for _, neighbor := range g.AdjList[current]{
			if !visited[neighbor]{
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	return nil // Path not found
}

// reconstructPath reconstructs the path from source to destination using the parent array
func (g *Graph) reconstructPath(parent []int, destination int) []int {
	path := []int(nil)
	
	for current := destination; current != -1; current = parent[current]{
		path = append([]int{current}, path...)
	}

	return path
}
