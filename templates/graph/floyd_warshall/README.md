# Floyd-Warshall Algorithm for All-Pairs Shortest Paths

## Problem Description

Implement the Floyd-Warshall algorithm to find the shortest paths between all pairs of vertices in a weighted graph. The algorithm should work with both positive and negative edge weights, but should detect negative cycles if they exist.

## Approach

The Floyd-Warshall algorithm uses dynamic programming to compute the shortest paths. It works by incrementally improving an estimate on the shortest path between two vertices, until the estimate is optimal.

The key steps of the Floyd-Warshall algorithm are:

1. Initialize the distance matrix with the direct edge weights between vertices.
2. For each intermediate vertex k:
   - For each pair of vertices (i, j):
     - If the path from i to j through k is shorter than the current path from i to j, update the shortest path.

## Complexity Analysis

- Time Complexity: O(V^3), where V is the number of vertices in the graph. The algorithm uses three nested loops, each iterating V times.
- Space Complexity: O(V^2) to store the distance matrix.

## Usage

```go
// Create a new graph with 4 vertices
g := NewGraph(4)

// Add edges to the graph
g.AddEdge(0, 1, 5)
g.AddEdge(0, 3, 10)
g.AddEdge(1, 2, 3)
g.AddEdge(2, 3, 1)

// Find all-pairs shortest paths
dist := g.FloydWarshall()

// Print the shortest distances
for i := 0; i < g.Vertices; i++ {
    for j := 0; j < g.Vertices; j++ {
        fmt.Printf("Shortest distance from %d to %d: %f\n", i, j, dist[i][j])
    }
}

// Get a specific path
path := g.GetPath(0, 3)
fmt.Println("Shortest path from 0 to 3:", path)

// Check for negative cycles
hasNegativeCycle := g.HasNegativeCycle()
fmt.Println("Graph has a negative cycle:", hasNegativeCycle)
```

## Implementation Details

The package provides the following main components:

1. `Graph`: Represents a weighted graph using an adjacency matrix.
2. `NewGraph`: Creates a new Graph with the given number of vertices.
3. `AddEdge`: Adds a weighted edge to the graph.
4. `FloydWarshall`: Runs the Floyd-Warshall algorithm to find all-pairs shortest paths.
5. `GetPath`: Retrieves the shortest path between two vertices.
6. `HasNegativeCycle`: Checks if the graph contains a negative cycle.

The `FloydWarshall` function returns a 2D slice representing the shortest distances between all pairs of vertices.

The `GetPath` function reconstructs the shortest path between two vertices using the results from the Floyd-Warshall algorithm.

The `HasNegativeCycle` function checks for the presence of negative cycles by examining the diagonal of the distance matrix after running the Floyd-Warshall algorithm.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Simple graphs with positive edge weights
2. Graphs with negative edge weights
3. Path reconstruction for existing and non-existing paths
4. Detection of negative cycles

To run the tests, use the following command:

```bash
go test
```

## Advantages of the Floyd-Warshall Algorithm

- Solves the all-pairs shortest path problem in a single execution
- Works with both positive and negative edge weights
- Can detect negative cycles in the graph
- Simple to implement and understand

## Limitations

- Not efficient for sparse graphs compared to running Dijkstra's algorithm for each vertex
- Requires O(V^2) memory, which can be prohibitive for very large graphs

## Applications

1. Finding the shortest paths in road networks or transportation systems
2. Network routing protocols
3. Detecting arbitrage opportunities in currency exchange
4. Calculating the transitive closure of a graph
5. Solving certain types of linear programming problems

The Floyd-Warshall algorithm is particularly useful in scenarios where you need to find shortest paths between all pairs of vertices in a graph, especially when the graph is dense or when negative edge weights are present. Its ability to detect negative cycles also makes it valuable in certain financial and optimization applications.
