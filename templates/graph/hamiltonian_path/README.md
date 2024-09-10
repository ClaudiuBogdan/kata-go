# Hamiltonian Path Finder

## Problem Description

Implement a function that finds a Hamiltonian path in an undirected graph if one exists. A Hamiltonian path is a path that visits each vertex exactly once.

## Approach

The implementation uses a backtracking algorithm to find a Hamiltonian path:

1. Start from each vertex as a potential starting point.
2. Add the next vertex to the path if it's adjacent to the previous vertex and hasn't been visited yet.
3. Recursively try to construct a Hamiltonian path using the remaining vertices.
4. If we can't construct a Hamiltonian path with the current vertex, we backtrack and try a different vertex.
5. If we've included all vertices in the path, we've found a Hamiltonian path.

## Complexity Analysis

- Time Complexity: O(n!), where n is the number of vertices. In the worst case, we try all permutations of vertices.
- Space Complexity: O(n) for the recursion stack and to store the path.

## Usage

```go
g := NewGraph(5)
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 3)
g.AddEdge(3, 4)

path := g.FindHamiltonianPath()
if path != nil {
    fmt.Println("Hamiltonian path found:", path)
} else {
    fmt.Println("No Hamiltonian path exists in this graph")
}
```

## Implementation Details

The package provides the following main components:

1. `Graph` struct: Represents an undirected graph using an adjacency list.
2. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
3. `AddEdge(v, w int)`: Adds an undirected edge between vertices v and w.
4. `FindHamiltonianPath() []int`: Finds a Hamiltonian path in the graph if one exists.

The implementation uses a backtracking algorithm with helper functions to check if adding a vertex to the current path is safe and to recursively construct the Hamiltonian path.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Graph with a Hamiltonian path
2. Graph without a Hamiltonian path
3. Complete graph
4. Single node graph
5. Two node graph

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the Hamiltonian path finder under different graph structures.

## Advantages and Limitations

Advantages:
- Correctly finds a Hamiltonian path if one exists
- Works for various graph structures
- Can start the path from any vertex in the graph

Limitations:
- Exponential time complexity, may be slow for large graphs
- Not suitable for real-time applications with large graphs

## Applications

- Solving puzzle games (e.g., Knight's tour problem)
- Genome assembly in bioinformatics
- Vehicle routing problems
- Optimizing computer networks

Remember that finding a Hamiltonian path is an NP-complete problem, and this implementation may not be efficient for very large graphs. For practical applications with large graphs, approximation algorithms or heuristics are often used instead.

## Difference from Hamiltonian Cycle

The main difference between this Hamiltonian path implementation and the previous Hamiltonian cycle implementation is that a Hamiltonian path doesn't need to return to the starting vertex. This allows for more flexibility in finding a valid path, as the start and end vertices can be different.

