# Maximum Bipartite Matching

## Problem Description

Implement an algorithm to find the maximum matching in a bipartite graph. A matching in a bipartite graph is a set of edges chosen in such a way that no two edges share an endpoint. A maximum matching is a matching that contains the largest possible number of edges.

## Approach

This implementation uses the Hungarian algorithm (also known as the Kuhn-Munkres algorithm) to find the maximum matching in a bipartite graph. The algorithm works as follows:

1. Start with an empty matching.
2. For each unmatched vertex in set U:
   a. Perform a depth-first search (DFS) to find an augmenting path.
   b. If an augmenting path is found, toggle the matching status of all edges on this path.
3. Repeat step 2 until no more augmenting paths can be found.

The DFS uses the following rules:
- If an unmatched vertex in V is found, the augmenting path ends here.
- If a matched vertex in V is found, continue the DFS from the corresponding matched vertex in U.

## Complexity Analysis

- Time Complexity: O(V * E), where V is the number of vertices and E is the number of edges. In the worst case, we might need to traverse all edges for each vertex in U.
- Space Complexity: O(V) for the visited array and the matching array.

## Usage

```go
g := NewBipartiteGraph(3, 3)
g.AddEdge(0, 0)
g.AddEdge(0, 1)
g.AddEdge(1, 0)
g.AddEdge(1, 1)
g.AddEdge(2, 1)
g.AddEdge(2, 2)

matching := g.FindMaximumMatching()
size := g.GetMatchingSize(matching)

fmt.Printf("Maximum matching size: %d\n", size)
fmt.Printf("Matching: %v\n", matching)
```

## Implementation Details

The package provides the following main components:

1. `BipartiteGraph` struct: Represents a bipartite graph with two sets of vertices U and V.
2. `NewBipartiteGraph(u, v int) *BipartiteGraph`: Creates a new BipartiteGraph with u vertices in set U and v vertices in set V.
3. `AddEdge(u, v int)`: Adds an edge from vertex u in set U to vertex v in set V.
4. `FindMaximumMatching() []int`: Finds the maximum matching in the bipartite graph.
5. `GetMatchingSize(matching []int) int`: Returns the size of the given matching.

The implementation uses an adjacency list to represent the graph and a depth-first search approach to find augmenting paths.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple bipartite graph
2. Complete bipartite graph
3. Empty graph
4. Unbalanced bipartite graph

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the maximum bipartite matching implementation under different graph structures.

## Advantages and Limitations

Advantages:
- Finds the maximum matching in any bipartite graph
- Works well for both sparse and dense graphs
- Can be extended to solve assignment problems and minimum cost flow problems

Limitations:
- May not be the most efficient for very large graphs (there are more advanced algorithms with better time complexity for dense graphs)
- Only works for bipartite graphs; cannot be used directly on general graphs

## Applications

- Assignment problems (e.g., assigning tasks to workers)
- Network flow problems
- Image processing (e.g., stereo matching in computer vision)
- Scheduling problems
- Recommendation systems

Remember that while this implementation finds a maximum matching, it may not find all possible maximum matchings if there are multiple solutions. If you need to enumerate all possible maximum matchings, you would need to modify the algorithm or use a different approach.

