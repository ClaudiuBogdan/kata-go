# Kruskal's Algorithm for Minimum Spanning Tree

## Problem Description

Implement Kruskal's algorithm to find the Minimum Spanning Tree (MST) of a weighted undirected graph. A minimum spanning tree is a subset of the edges that connects all vertices with the minimum total edge weight, without creating any cycles.

## Approach

Kruskal's algorithm follows a greedy approach to find the MST:

1. Sort all edges in non-decreasing order of their weight.
2. Initialize a disjoint set data structure for all vertices.
3. Iterate through the sorted edges:
   a. If including the current edge doesn't form a cycle (i.e., the vertices are in different sets), add it to the MST.
   b. Union the sets of the two vertices connected by the edge.
4. Repeat step 3 until we have V-1 edges in the MST (where V is the number of vertices).

## Complexity Analysis

- Time Complexity: O(E log E) or O(E log V), where E is the number of edges and V is the number of vertices. This is due to the sorting of edges.
- Space Complexity: O(V + E) for storing the graph and the disjoint set data structure.

## Usage

```go
g := NewGraph(4)
g.AddEdge(0, 1, 1)
g.AddEdge(0, 2, 4)
g.AddEdge(0, 3, 3)
g.AddEdge(1, 2, 2)
g.AddEdge(2, 3, 5)

mst := g.KruskalMST()
fmt.Println("Minimum Spanning Tree edges:")
for _, edge := range mst {
    fmt.Printf("(%d, %d) with weight %d\n", edge.Src, edge.Dest, edge.Weight)
}
```

## Implementation Details

The package provides the following main components:

1. `Edge` struct: Represents an edge in the graph with source, destination, and weight.
2. `Graph` struct: Represents a weighted undirected graph using a list of edges.
3. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
4. `AddEdge(src, dest, weight int)`: Adds a weighted undirected edge to the graph.
5. `KruskalMST() []Edge`: Finds the Minimum Spanning Tree using Kruskal's algorithm.

The implementation uses a disjoint set data structure (implemented with the `find` and `union` functions) to efficiently detect cycles while building the MST.

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

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of Kruskal's algorithm implementation under different graph structures.

## Advantages and Limitations

Advantages:

- Guarantees the minimum spanning tree for a connected, undirected graph
- Works well for sparse graphs
- Easy to implement and understand

Limitations:

- May not be the most efficient for dense graphs (Prim's algorithm might be better in such cases)
- Requires sorting all edges, which can be memory-intensive for very large graphs

## Applications

- Network design (e.g., laying out electrical wiring, computer networks, or telecommunications networks)
- Approximation algorithms for NP-hard problems (e.g., traveling salesman problem)
- Cluster analysis in data mining
- Image segmentation in computer vision

Remember that Kruskal's algorithm is just one way to find a minimum spanning tree. For dense graphs or when starting from a specific vertex is preferred, Prim's algorithm might be a better choice.

##

[MST Explanation](https://www.youtube.com/watch?v=71UQH7Pr9kU)
