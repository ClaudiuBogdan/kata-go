# Minimum Cut in a Graph

## Problem Description

Implement an algorithm to find the minimum cut in an undirected graph. A cut in a graph is a partition of the vertices into two disjoint subsets. The cut-set of a cut is the set of edges whose endpoints are in different subsets of the partition. The value of a cut is the sum of the weights of the edges in the cut-set. The minimum cut is the cut with the smallest value among all possible cuts.

## Approach

This implementation uses the Stoer-Wagner algorithm to find the global minimum cut in an undirected graph. The algorithm works as follows:

1. Start with the original graph.
2. Repeat n-1 times (where n is the number of vertices):
   a. Find a minimum s-t cut using the maximum adjacency search.
   b. If this cut is smaller than the best cut found so far, update the best cut.
   c. Merge the two vertices s and t.
3. Return the best cut found.

## Complexity Analysis

- Time Complexity: O(V^3) or O(V*E*log(V)) depending on the implementation of the maximum adjacency search. This implementation uses a simple O(V^2) approach for the search, resulting in an overall O(V^3) time complexity.
- Space Complexity: O(V^2) for storing the graph as an adjacency matrix.

## Usage

```go
g := NewGraph(4)
g.AddEdge(0, 1, 2)
g.AddEdge(0, 2, 3)
g.AddEdge(0, 3, 3)
g.AddEdge(1, 2, 1)
g.AddEdge(2, 3, 1)

minCut, partition := g.FindMinCut()
fmt.Printf("Minimum cut value: %d\n", minCut)
fmt.Printf("Partition: %v\n", partition)
```

## Implementation Details

The package provides the following main components:

1. `Graph` struct: Represents an undirected graph using an adjacency matrix.
2. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
3. `AddEdge(u, v, w int)`: Adds an undirected edge between vertices u and v with weight w.
4. `FindMinCut() (int, []int)`: Finds the minimum cut in the graph using the Stoer-Wagner algorithm, returning the cut value and one side of the partition.

The implementation uses an adjacency matrix to represent the graph for simplicity and ease of understanding. For very large sparse graphs, an adjacency list representation might be more memory-efficient, but would require modifications to the algorithm implementation.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple graph
2. Complex graph
3. Disconnected graph
4. Complete graph

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the minimum cut implementation under different graph structures.

## Advantages and Limitations

Advantages:
- Finds the global minimum cut in any undirected graph
- Relatively simple to understand and implement
- Works for both weighted and unweighted graphs

Limitations:
- May not be the most efficient for very large graphs (there are more advanced algorithms with better time complexity for sparse graphs)
- The adjacency matrix representation may consume too much memory for extremely large sparse graphs
- Only works for undirected graphs; for directed graphs, different algorithms are needed

## Applications

- Network reliability analysis
- Image segmentation in computer vision
- Clustering in data mining
- Social network analysis
- Circuit partitioning in VLSI design

Remember that while this implementation finds the minimum cut, it returns only one side of the partition. The other side can be easily derived by taking the complement of the returned partition with respect to all vertices in the graph.

