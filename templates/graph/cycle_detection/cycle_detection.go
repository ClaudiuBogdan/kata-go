package cycledetection

// Graph represents a directed graph using an adjacency list
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

// AddEdge adds a directed edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.AdjList[v] = append(g.AdjList[v], w)
}

// DetectCycle checks if the graph contains a cycle
func (g *Graph) DetectCycle() bool {
	visited := make([]bool, g.Vertices)
	recStack := make([]bool, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		if g.dfsDetectCycle(i, visited, recStack) {
			return true
		}
	}

	return false
}

// dfsDetectCycle performs a depth-first search to detect cycles
func (g *Graph) dfsDetectCycle(v int, visited, recStack []bool) bool {
	if !visited[v] {
		visited[v] = true
		recStack[v] = true

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] && g.dfsDetectCycle(neighbor, visited, recStack) {
				return true
			} else if recStack[neighbor] {
				return true
			}
		}
	}

	recStack[v] = false
	return false
}

// FindCycle returns a cycle in the graph if one exists, otherwise returns nil
func (g *Graph) FindCycle() []int {
	visited := make([]bool, g.Vertices)
	recStack := make([]bool, g.Vertices)
	parent := make([]int, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		parent[i] = -1
	}

	for i := 0; i < g.Vertices; i++ {
		if cycle := g.dfsFindCycle(i, visited, recStack, parent); cycle != nil {
			return cycle
		}
	}

	return nil
}

// dfsFindCycle performs a depth-first search to find a cycle
func (g *Graph) dfsFindCycle(v int, visited, recStack []bool, parent []int) []int {
	if !visited[v] {
		visited[v] = true
		recStack[v] = true

		for _, neighbor := range g.AdjList[v] {
			if !visited[neighbor] {
				parent[neighbor] = v
				if cycle := g.dfsFindCycle(neighbor, visited, recStack, parent); cycle != nil {
					return cycle
				}
			} else if recStack[neighbor] {
				return g.constructCycle(v, neighbor, parent)
			}
		}
	}

	recStack[v] = false
	return nil
}

// constructCycle builds the cycle path from the detected cycle
func (g *Graph) constructCycle(start, end int, parent []int) []int {
	cycle := []int{end, start}
	for v := parent[start]; v != end; v = parent[v] {
		cycle = append([]int{v}, cycle...)
	}
	return cycle
}
