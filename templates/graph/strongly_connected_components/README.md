# Strongly Connected Components in a Directed Graph

## Problem Description

Implement an algorithm to find strongly connected components (SCCs) in a directed graph. A strongly connected component is a maximal subgraph of a directed graph in which there is a path from each vertex to every other vertex.

## Approach

This implementation uses Kosaraju's algorithm to find strongly connected components. The algorithm works in two main steps:

1. Perform a depth-first search (DFS) on the original graph, keeping track of the finish times of vertices.
2. Create a transposed graph (reverse all edge directions).
3. Perform a second DFS on the transposed graph, starting with vertices in decreasing order of finish time.

Each tree in the second DFS forest is a strongly connected component.

## Complexity Analysis

- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges.
- Space Complexity: O(V) for the visited array, stack, and output.

## Usage

```go
g := NewGraph(5)
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 0)
g.AddEdge(1, 3)
g.AddEdge(3, 4)

sccs := g.FindSCCs()
fmt.Println("Strongly Connected Components:")
for i, scc := range sccs {
    fmt.Printf("SCC %d: %v\n", i+1, scc)
}
```

## Implementation Details

The package provides the following main components:

1. `Graph` struct: Represents a directed graph using an adjacency list.
2. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
3. `AddEdge(u, v int)`: Adds a directed edge from u to v.
4. `FindSCCs() [][]int`: Finds the strongly connected components in the graph.

The implementation uses depth-first search (DFS) and graph transposition to efficiently find SCCs.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple graph
2. Complex graph
3. Single vertex graph
4. Disconnected graph

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the strongly connected components implementation under different graph structures.

## Advantages and Limitations

Advantages:
- Efficiently finds all strongly connected components in a directed graph
- Works well for both sparse and dense graphs
- Linear time complexity O(V + E)

Limitations:
- Requires two passes through the graph
- May not be the most intuitive algorithm to understand at first glance

## Applications

- Analyzing the structure of social networks
- Detecting cycles in directed graphs
- Solving 2-satisfiability problems
- Analyzing dependencies in software packages
- Identifying strongly linked pages in web graphs

Remember that Kosaraju's algorithm is just one way to find strongly connected components. For very large graphs or when memory is a concern, Tarjan's algorithm might be a better choice as it requires only one pass through the graph.

