# Finding Eulerian Circuits in a Graph

## Problem Description

Implement an algorithm to find an Eulerian circuit in an undirected graph. An Eulerian circuit is a path in a graph that visits every edge exactly once and returns to the starting vertex. The algorithm should determine if an Eulerian circuit exists and, if so, find one.

## Approach

The algorithm uses the following approach:

1. Check if the graph has an Eulerian circuit:
   - The graph must be connected (except for isolated vertices).
   - All vertices must have an even degree.

2. If an Eulerian circuit exists, find it using a depth-first search (DFS) approach:
   - Start from any vertex.
   - Follow unused edges, removing them from the graph as they're used.
   - When stuck, add the current vertex to the circuit and backtrack.

This approach is known as Hierholzer's algorithm.

## Complexity Analysis

- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges. We visit each vertex and edge once during the DFS.
- Space Complexity: O(V + E) to store the graph and the resulting circuit.

## Usage

```go
// Create a new graph with 3 vertices
g := NewGraph(3)

// Add edges to create a simple cycle
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 0)

// Check if the graph has an Eulerian circuit
hasCircuit := g.HasEulerianCircuit()
fmt.Println("Graph has an Eulerian circuit:", hasCircuit)

// Find an Eulerian circuit
circuit := g.FindEulerianCircuit()
if circuit != nil {
    fmt.Println("Eulerian circuit:", circuit)
} else {
    fmt.Println("No Eulerian circuit exists")
}
```

## Implementation Details

The package provides the following main components:

1. `Graph`: Represents an undirected graph using an adjacency list.
2. `NewGraph`: Creates a new Graph with the given number of vertices.
3. `AddEdge`: Adds an undirected edge to the graph.
4. `HasEulerianCircuit`: Checks if the graph has an Eulerian circuit.
5. `FindEulerianCircuit`: Finds an Eulerian circuit in the graph if one exists.

The `HasEulerianCircuit` function checks two conditions:
1. The graph is connected (using a depth-first search).
2. All vertices have an even degree.

The `FindEulerianCircuit` function implements Hierholzer's algorithm to find an Eulerian circuit. It uses a depth-first search approach, removing edges as they are traversed and backtracking when necessary.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Graphs with Eulerian circuits
2. Graphs without Eulerian circuits (odd degree vertices)
3. Disconnected graphs
4. Empty graphs
5. Complex graphs with multiple cycles

The tests also include a validation function to ensure that the found circuit is indeed a valid Eulerian circuit.

To run the tests, use the following command:

```bash
go test
```

## Advantages of Finding Eulerian Circuits

- Solving "closed path" problems in various applications
- Optimizing routes in logistics and transportation
- Analyzing and designing circuits in computer networking

## Applications

1. DNA fragment assembly in bioinformatics
2. Optimizing snow plow routes in city planning
3. Solving maze puzzles
4. Designing efficient network protocols
5. Optimizing pen plotters in computer graphics

## Limitations and Considerations

- This implementation is for undirected graphs. For directed graphs, the concept of Eulerian circuits is slightly different (all vertices must have equal in-degree and out-degree).
- The algorithm assumes that the graph is represented using an adjacency list. For very dense graphs, an adjacency matrix representation might be more efficient.
- For graphs with a large number of vertices but relatively few edges, more efficient data structures (like Fibonacci heaps) can be used to improve performance.

Finding Eulerian circuits is a classic problem in graph theory with numerous practical applications. By efficiently determining the existence of and finding Eulerian circuits, we can solve a wide range of problems in various domains, from route optimization to DNA sequencing.
