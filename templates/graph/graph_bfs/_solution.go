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
	g.AdjList[v] = append(g.AdjList[v], w)
	g.AdjList[w] = append(g.AdjList[w], v)
}

// BFS performs a breadth-first search starting from the given source vertex
func (g *Graph) BFS(source int) []int {
	visited := make([]bool, g.Vertices)
	queue := []int{source}
	visited[source] = true

	result := []int{}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return result
}

// ShortestPath finds the shortest path between source and destination using BFS
func (g *Graph) ShortestPath(source, destination int) []int {
	visited := make([]bool, g.Vertices)
	parent := make([]int, g.Vertices)
	for i := range parent {
		parent[i] = -1
	}

	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v == destination {
			return g.reconstructPath(parent, source, destination)
		}

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = v
			}
		}
	}

	return nil // No path found
}

// reconstructPath reconstructs the path from source to destination using the parent array
func (g *Graph) reconstructPath(parent []int, source, destination int) []int {
	path := []int{}
	for v := destination; v != -1; v = parent[v] {
		path = append([]int{v}, path...)
	}
	return path
}
