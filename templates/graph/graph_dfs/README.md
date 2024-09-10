# Depth-First Search (DFS) on a Graph

## Problem Description

Implement a Depth-First Search (DFS) algorithm on an undirected graph represented as an adjacency list. The algorithm should traverse the graph starting from a given source vertex, check if a path exists between two vertices, and find all possible paths between two vertices.

## Approach

DFS explores as far as possible along each branch before backtracking. The algorithm uses recursion (or a stack for an iterative approach) to keep track of the vertices to be explored.

The key steps of the DFS algorithm are:

1. Start from the given source vertex and mark it as visited.
2. Explore an unvisited adjacent vertex of the current vertex.
3. Repeat step 2 until there are no more unvisited adjacent vertices.
4. Backtrack to the previous vertex and repeat steps 2-3.
5. Repeat steps 2-4 until all vertices are visited.

## Complexity Analysis

- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges in the graph. In the worst case, we might need to visit all vertices and edges.
- Space Complexity: O(V) for the recursion stack (or explicit stack in iterative implementation) and visited array.

## Usage

```go
// Create a new graph with 5 vertices
g := NewGraph(5)

// Add edges to the graph
g.AddEdge(0, 1)
g.AddEdge(0, 2)
g.AddEdge(1, 3)
g.AddEdge(2, 3)
g.AddEdge(3, 4)

// Perform DFS starting from vertex 0
dfsOrder := g.DFS(0)
fmt.Println("DFS order:", dfsOrder)

// Check if a path exists between 0 and 4
hasPath := g.HasPath(0, 4)
fmt.Println("Path exists between 0 and 4:", hasPath)

// Find all paths between 0 and 4
allPaths := g.FindAllPaths(0, 4)
fmt.Println("All paths between 0 and 4:", allPaths)
```

## Implementation Details

The package provides the following main components:

1. `Graph`: Represents an undirected graph using an adjacency list.
2. `NewGraph`: Creates a new Graph with the given number of vertices.
3. `AddEdge`: Adds an undirected edge to the graph.
4. `DFS`: Performs a depth-first search starting from the given source vertex.
5. `HasPath`: Checks if a path exists between two vertices using DFS.
6. `FindAllPaths`: Finds all possible paths between two vertices using DFS.

The `DFS` function returns a slice of integers representing the order in which vertices are visited during the DFS traversal.

The `HasPath` function returns a boolean indicating whether a path exists between the source and destination vertices.

The `FindAllPaths` function returns a slice of slices, where each inner slice represents a path from the source to the destination vertex.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. DFS traversal on different graph structures (connected, disconnected, linear)
2. Path existence checking for various cases
3. Finding all paths between two vertices

To run the tests, use the following command:

```bash
go test
```

## Advantages of DFS

- Uses less memory than BFS for very deep graphs
- Useful for tasks such as topological sorting, finding strongly connected components, and maze solving
- Can be easily implemented using recursion

## Applications

1. Topological sorting in directed acyclic graphs
2. Finding strongly connected components in a graph
3. Maze solving and puzzle games
4. Cycle detection in graphs
5. Web crawling
6. Analyzing game states in AI (e.g., chess game tree analysis)
7. Generating permutations and combinations

## Limitations and Considerations

- Not guaranteed to find the shortest path in unweighted graphs (use BFS for that)
- Can get stuck in infinite loops if not implemented carefully in graphs with cycles
- The recursive implementation might cause stack overflow for very deep graphs (consider an iterative approach for such cases)
- The order of visiting neighbors can affect the result, so it may not always produce the same output for graphs with multiple valid DFS orderings

Depth-First Search is a fundamental graph traversal algorithm with numerous applications in computer science and beyond. Its ability to deeply explore paths before backtracking makes it particularly useful for problems involving path finding, cycle detection, and exhaustive search scenarios.
