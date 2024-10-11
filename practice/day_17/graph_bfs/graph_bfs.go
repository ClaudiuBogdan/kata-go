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
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {
	// TODO: add edge in both ways, as the graph is undirected
	g.AdjList[v] = append(g.AdjList[v], w)
	g.AdjList[w] = append(g.AdjList[w], v)
}

// BFS performs a breadth-first search starting from the given source vertex
func (g *Graph) BFS(source int) []int {
	visited := make(map[int]bool)
	queue := []int{source}
	result := []int(nil)
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range g.AdjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// ShortestPath finds the shortest path between source and destination using BFS
func (g *Graph) ShortestPath(source, destination int) []int {
	parent := make([]int, g.Vertices)
	visited := make(map[int]bool)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == destination {
			return g.reconstructPath(parent, source, destination)
		}

		for _, neighbor := range g.AdjList[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				parent[neighbor] = current
				visited[neighbor] = true
			}
		}
	}

	return nil
}

// reconstructPath reconstructs the path from source to destination using the parent array
func (g *Graph) reconstructPath(parent []int, source, destination int) []int {
	path := []int{destination}
	for current := destination; current != source; current = parent[current] {
		path = append([]int{parent[current]}, path...)
	}
	return path
}
