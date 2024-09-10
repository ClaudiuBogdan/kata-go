# Bellman-Ford Algorithm for Shortest Paths

## Problem Description

Implement the Bellman-Ford algorithm to find the shortest paths from a source vertex to all other vertices in a weighted graph. The algorithm should also detect negative cycles in the graph.

## Approach

The Bellman-Ford algorithm works by relaxing all edges V-1 times, where V is the number of vertices in the graph. After this process, it checks for negative cycles by attempting one more relaxation. If any distance can be further reduced, it indicates the presence of a negative cycle.

The key steps of the Bellman-Ford algorithm are:

1. Initialize distances from the source to all vertices as infinite and distance to the source itself as 0.
2. Repeat V-1 times:
   - For each edge (u, v) with weight w, if dist[v] > dist[u] + w, then update dist[v] = dist[u] + w.
3. Check for negative-weight cycles:
   - For each edge (u, v) with weight w, if dist[v] > dist[u] + w, then the graph contains a negative-weight cycle.

## Complexity Analysis

- Time Complexity: O(VE), where V is the number of vertices and E is the number of edges in the graph.
- Space Complexity: O(V), used to store the distance array.

## Usage

```go
// Create a new graph with 5 vertices
g := NewGraph(5)

// Add edges to the graph
g.AddEdge(0, 1, -1)
g.AddEdge(0, 2, 4)
g.AddEdge(1, 2, 3)
g.AddEdge(1, 3, 2)
g.AddEdge(1, 4, 2)
g.AddEdge(3, 2, 5)
g.AddEdge(3, 1, 1)
g.AddEdge(4, 3, -3)

// Find shortest paths from source vertex 0
distances, hasNegativeCycle := g.BellmanFord(0)

if hasNegativeCycle {
    fmt.Println("Graph contains a negative cycle")
} else {
    fmt.Println("Shortest distances from source:", distances)
}

// Get the shortest path from 0 to 3
path := g.GetPath(0, 3)
fmt.Println("Shortest path from 0 to 3:", path)
```

## Implementation Details

The package provides the following main components:

1. `Edge`: Represents a weighted edge in the graph.
2. `Graph`: Represents a weighted graph using a list of edges.
3. `NewGraph`: Creates a new Graph with the given number of vertices.
4. `AddEdge`: Adds a weighted edge to the graph.
5. `BellmanFord`: Runs the Bellman-Ford algorithm to find the shortest paths from a source vertex.
6. `GetPath`: Retrieves the shortest path from the source to a destination vertex.

The `BellmanFord` function returns two values: an array of shortest distances from the source to all vertices, and a boolean indicating whether a negative cycle was detected. If a negative cycle is present, the distance array will be nil.

The `GetPath` function reconstructs the shortest path from the source to a given destination using the results from the Bellman-Ford algorithm.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Simple graph with positive and negative edge weights
2. Graph with a negative cycle
3. Disconnected graph
4. Path reconstruction for various cases

To run the tests, use the following command:

```bash
go test
```

## Advantages of the Bellman-Ford Algorithm

- Can handle graphs with negative edge weights (unlike Dijkstra's algorithm)
- Detects negative cycles in the graph
- Simple to implement and understand

## Limitations

- Slower than Dijkstra's algorithm for graphs with non-negative edge weights
- Not suitable for graphs with a large number of edges due to its time complexity

## Applications

- Routing protocols in computer networks (e.g., RIP - Routing Information Protocol)
- Arbitrage detection in financial markets
- Traffic routing systems
- Solving constraint satisfaction problems in artificial intelligence

The Bellman-Ford algorithm is particularly useful in scenarios where negative edge weights are possible or when it's necessary to detect negative cycles in the graph. Its ability to handle these cases makes it a versatile tool in various domains, despite its higher time complexity compared to some other shortest path algorithms.
