# Cycle Detection in a Graph

## Problem Description

Implement an algorithm to detect cycles in a directed graph. The algorithm should be able to determine whether a cycle exists in the graph and, optionally, find and return a specific cycle if one exists.

## Approach

The algorithm uses a depth-first search (DFS) traversal to detect cycles in the graph. It maintains two arrays:

1. `visited`: Keeps track of vertices that have been visited during the DFS traversal.
2. `recStack`: Keeps track of vertices currently in the recursion stack of the DFS.

The key idea is that if we encounter a vertex that is already in the recursion stack, we have found a cycle.

For cycle finding, we also maintain a `parent` array to reconstruct the cycle path once detected.

## Complexity Analysis

- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges in the graph. We perform a DFS traversal of the graph, visiting each vertex and edge once.
- Space Complexity: O(V), used for the visited and recursion stack arrays, and the adjacency list representation of the graph.

## Usage

```go
// Create a new graph with 4 vertices
g := NewGraph(4)

// Add edges to create a cycle
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 3)
g.AddEdge(3, 1)

// Detect if there's a cycle
hasCycle := g.DetectCycle()
fmt.Println("Graph has a cycle:", hasCycle)

// Find a specific cycle
cycle := g.FindCycle()
if cycle != nil {
    fmt.Println("Cycle found:", cycle)
} else {
    fmt.Println("No cycle found")
}
```

## Implementation Details

The package provides the following main components:

1. `Graph`: Represents a directed graph using an adjacency list.
2. `NewGraph`: Creates a new Graph with the given number of vertices.
3. `AddEdge`: Adds a directed edge to the graph.
4. `DetectCycle`: Checks if the graph contains a cycle.
5. `FindCycle`: Returns a specific cycle in the graph if one exists.

The `DetectCycle` function returns a boolean indicating whether a cycle was found in the graph.

The `FindCycle` function returns a slice of integers representing the vertices in a cycle, or nil if no cycle is found.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Cyclic graphs
2. Acyclic graphs
3. Graphs with self-loops
4. Disconnected graphs
5. Empty graphs
6. Graphs with multiple cycles

To run the tests, use the following command:

```bash
go test
```

## Advantages of Cycle Detection

- Identifying circular dependencies in systems
- Detecting deadlocks in resource allocation
- Ensuring the correctness of certain graph algorithms that require acyclic graphs

## Applications

1. Dependency resolution in build systems and package managers
2. Deadlock detection in operating systems and distributed systems
3. Cycle detection in financial transaction systems to prevent fraud
4. Analyzing workflows and business processes for inefficiencies
5. Detecting infinite loops in program analysis

## Limitations and Considerations

- This implementation is for directed graphs. For undirected graphs, the algorithm would need to be modified to handle bidirectional edges.
- The current implementation assumes that vertex indices are continuous from 0 to n-1. For graphs with non-continuous vertex identifiers, the data structure would need to be adapted.
- For very large graphs, a non-recursive implementation might be preferred to avoid potential stack overflow issues.

Cycle detection is a fundamental graph algorithm with wide-ranging applications in computer science and beyond. By identifying cycles in graphs, we can solve various problems related to dependency management, resource allocation, and system analysis.
