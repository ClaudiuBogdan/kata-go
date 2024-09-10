# Hamiltonian Cycle Finder

## Problem Description

Implement a function that finds a Hamiltonian cycle in an undirected graph if one exists. A Hamiltonian cycle is a cycle that visits each vertex exactly once and returns to the starting vertex.

## Approach

The implementation uses a backtracking algorithm to find a Hamiltonian cycle:

1. Start from the first vertex (index 0).
2. Add the next vertex to the path if it's adjacent to the previous vertex and hasn't been visited yet.
3. Recursively try to construct a Hamiltonian cycle using the remaining vertices.
4. If we can't construct a Hamiltonian cycle with the current vertex, we backtrack and try a different vertex.
5. If we've included all vertices in the path and there's an edge from the last vertex to the first vertex, we've found a Hamiltonian cycle.

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
g.AddEdge(4, 0)
g.AddEdge(0, 3)
g.AddEdge(1, 4)

cycle := g.FindHamiltonianCycle()
if cycle != nil {
    fmt.Println("Hamiltonian cycle found:", cycle)
} else {
    fmt.Println("No Hamiltonian cycle exists in this graph")
}
```

## Implementation Details

The package provides the following main components:

1. `Graph` struct: Represents an undirected graph using an adjacency list.
2. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
3. `AddEdge(v, w int)`: Adds an undirected edge between vertices v and w.
4. `FindHamiltonianCycle() []int`: Finds a Hamiltonian cycle in the graph if one exists.

The implementation uses a backtracking algorithm with helper functions to check if adding a vertex to the current path is safe and to recursively construct the Hamiltonian cycle.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Graph with a Hamiltonian cycle
2. Graph without a Hamiltonian cycle
3. Complete graph
4. Single node graph

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the Hamiltonian cycle finder under different graph structures.

## Advantages and Limitations

Advantages:
- Correctly finds a Hamiltonian cycle if one exists
- Works for various graph structures

Limitations:
- Exponential time complexity, may be slow for large graphs
- Not suitable for real-time applications with large graphs

## Applications

- Solving the Traveling Salesman Problem (TSP)
- Circuit board design
- Network routing protocols
- DNA fragment assembly in bioinformatics

Remember that finding a Hamiltonian cycle is an NP-complete problem, and this implementation may not be efficient for very large graphs. For practical applications with large graphs, approximation algorithms or heuristics are often used instead.
