# Prim's Algorithm for Minimum Spanning Tree

## Problem Description

Implement Prim's algorithm to find the Minimum Spanning Tree (MST) of a weighted undirected graph. A minimum spanning tree is a subset of the edges that connects all vertices with the minimum total edge weight, without creating any cycles.

## Approach

This implementation uses Prim's algorithm with a priority queue to efficiently find the MST. The algorithm works as follows:

1. Start with an arbitrary vertex and mark it as visited.
2. Add all edges connected to the visited vertex to a priority queue.
3. While the priority queue is not empty and we haven't added V-1 edges to the MST:
   a. Extract the minimum weight edge from the priority queue.
   b. If the edge connects to an unvisited vertex:
      - Mark the vertex as visited.
      - Add the edge to the MST.
      - Add all edges connected to this new vertex to the priority queue.
4. The resulting set of edges forms the MST.

## Complexity Analysis

- Time Complexity: O((V + E) log V), where V is the number of vertices and E is the number of edges. This is achieved by using a binary heap as the priority queue.
- Space Complexity: O(V + E) for storing the graph as an adjacency list and the priority queue.

## Usage

```go
g := NewGraph(5)
g.AddEdge(0, 1, 2)
g.AddEdge(0, 3, 6)
g.AddEdge(1, 2, 3)
g.AddEdge(1, 3, 8)
g.AddEdge(1, 4, 5)
g.AddEdge(2, 4, 7)
g.AddEdge(3, 4, 9)

mst := g.PrimMST()
fmt.Println("Minimum Spanning Tree edges:")
for _, edge := range mst {
    fmt.Printf("(%d, %d) with weight %d\n", edge.From, edge.To, edge.Weight)
}
```

## Implementation Details

The package provides the following main components:

1. `Edge` struct: Represents an edge in the graph with destination vertex and weight.
2. `Graph` struct: Represents an undirected graph using an adjacency list.
3. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
4. `AddEdge(u, v, w int)`: Adds an undirected edge between vertices u and v with weight w.
5. `PrimMST() []Edge`: Finds the Minimum Spanning Tree using Prim's algorithm.

The implementation uses Go's `container/heap` package to implement the priority queue, which is crucial for the efficiency of Prim's algorithm.

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

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of Prim's algorithm implementation under different graph structures.

## Advantages and Limitations

Advantages:
- Efficiently finds the minimum spanning tree for connected, undirected graphs
- Works well for dense graphs
- Easy to implement with a priority queue

Limitations:
- May not be the most efficient for very sparse graphs (Kruskal's algorithm might be better in such cases)
- Requires a connected graph to find a spanning tree (will only find a spanning tree for one component in disconnected graphs)

## Applications

- Network design (e.g., laying out electrical wiring, computer networks, or telecommunications networks)
- Approximation algorithms for NP-hard problems (e.g., traveling salesman problem)
- Cluster analysis in data mining
- Image segmentation in computer vision
- Pathfinding algorithms in AI and robotics

Remember that Prim's algorithm is just one way to find a minimum spanning tree. For very sparse graphs or when the edges are already sorted, Kruskal's algorithm might be a better choice.

