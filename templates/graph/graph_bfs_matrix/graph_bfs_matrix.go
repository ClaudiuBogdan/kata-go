package bfsmatrix

// Graph represents an undirected graph using an adjacency matrix
type Graph struct {
	Vertices int
	Matrix   [][]int
}

// NewGraph creates a new Graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &Graph{
		Vertices: vertices,
		Matrix:   matrix,
	}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.Matrix[v][w] = 1
	g.Matrix[w][v] = 1
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

		for i := 0; i < g.Vertices; i++ {
			if g.Matrix[v][i] == 1 && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
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

		for i := 0; i < g.Vertices; i++ {
			if g.Matrix[v][i] == 1 && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
				parent[i] = v
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
