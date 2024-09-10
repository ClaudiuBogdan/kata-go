# Dijkstra's Algorithm for Shortest Paths

## Problem Description

Implement Dijkstra's algorithm to find the shortest paths from a source vertex to all other vertices in a weighted graph. The implementation should use an adjacency list representation for the graph and a priority queue for efficient selection of the next vertex to process.

## Approach

Dijkstra's algorithm works by maintaining a set of vertices whose shortest distance from the source is known. In each step, it selects the vertex with the minimum distance from the set of unprocessed vertices and updates the distances to its adjacent vertices.

The key steps of Dijkstra's algorithm are:

1. Initialize distances from the source to all vertices as infinite and distance to the source itself as 0.
2. Create a priority queue and enqueue the source vertex with distance 0.
3. While the priority queue is not empty:
   - Dequeue a vertex u with the minimum distance.
   - For each adjacent vertex v of u:
     - If the distance to v through u is less than its current distance, update its distance and enqueue it with the new distance.

## Complexity Analysis

- Time Complexity: O((V + E) log V), where V is the number of vertices and E is the number of edges in the graph. This is achieved by using a binary heap as the priority queue.
- Space Complexity: O(V + E), used for the adjacency list representation of the graph and the priority queue.

## Usage

```go
// Create a new graph with 5 vertices
g := NewGraph(5)

// Add edges to the graph
g.AddEdge(0, 1, 4)
g.AddEdge(0, 2, 1)
g.AddEdge(1, 3, 1)
g.AddEdge(2, 1, 2)
g.AddEdge(2, 3, 5)
g.AddEdge(3, 4, 3)

// Find shortest paths from source vertex 0
distances, predecessors := g.Dijkstra(0)

fmt.Println("Shortest distances from source:", distances)
fmt.Println("Predecessors:", predecessors)

// Get the shortest path from 0 to 4
path := g.GetPath(0, 4)
fmt.Println("Shortest path from 0 to 4:", path)
```

## Implementation Details

The package provides the following main components:

1. `Edge`: Represents a weighted edge in the graph.
2. `Graph`: Represents a weighted graph using an adjacency list.
3. `NewGraph`: Creates a new Graph with the given number of vertices.
4. `AddEdge`: Adds a weighted edge to the graph.
5. `Dijkstra`: Runs Dijkstra's algorithm to find the shortest paths from a source vertex.
6. `GetPath`: Retrieves the shortest path from the source to a destination vertex.

The `Dijkstra` function returns two slices: `distances` containing the shortest distances from the source to all vertices, and `predecessors` containing the predecessor of each vertex in the shortest path tree.

The `GetPath` function reconstructs the shortest path from the source to a given destination using the results from Dijkstra's algorithm.

The implementation uses a priority queue (implemented as a binary heap) for efficient selection of the next vertex to process.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Simple connected graph
2. Disconnected graph
3. Graph with no edges
4. Path finding for existing paths
5. Path finding when no path exists
6. Path finding when source is the destination

To run the tests, use the following command:

```bash
go test
```

## Advantages of Dijkstra's Algorithm

- Finds the shortest path in weighted graphs with non-negative edge weights
- Efficient for sparse graphs when using a binary heap
- Can be used to find shortest paths from a single source to all other vertices

## Limitations

- Does not work with negative edge weights (use the Bellman-Ford algorithm for such cases)
- May be less efficient for dense graphs compared to the Floyd-Warshall algorithm

## Applications

1. Network routing protocols (e.g., OSPF - Open Shortest Path First)
2. GPS navigation systems
3. Solving puzzles and games (e.g., finding the shortest path in a maze)
4. Social network analysis (finding the shortest connection between users)
5. Robotics (path planning for robots)

Dijkstra's algorithm is a fundamental graph algorithm with wide-ranging applications in computer science and beyond. Its ability to efficiently find shortest paths in weighted graphs makes it a crucial tool in various domains, from network routing to artificial intelligence.
