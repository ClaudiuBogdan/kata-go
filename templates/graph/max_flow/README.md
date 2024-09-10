# Maximum Flow in a Flow Network

## Problem Description

Implement an algorithm to find the maximum flow in a flow network. A flow network is a directed graph where each edge has a capacity, and we want to find the maximum flow from a source vertex to a sink vertex while respecting the capacity constraints.

## Approach

This implementation uses the Ford-Fulkerson algorithm with the Edmonds-Karp implementation for finding augmenting paths. The algorithm works as follows:

1. Start with zero flow.
2. While there exists an augmenting path from source to sink:
   a. Find an augmenting path using Breadth-First Search (BFS).
   b. Determine the bottleneck capacity along this path.
   c. Augment the flow along this path by the bottleneck capacity.
3. When no augmenting path can be found, the algorithm terminates, and the current flow is the maximum flow.

## Complexity Analysis

- Time Complexity: O(V * E^2), where V is the number of vertices and E is the number of edges. This is because we can have at most O(VE) augmenting paths, and each BFS takes O(E) time.
- Space Complexity: O(V^2) for storing the graph and flow matrices, and O(V) for the BFS queue and parent array.

## Usage

```go
fn := NewFlowNetwork(6)
fn.AddEdge(0, 1, 16)
fn.AddEdge(0, 2, 13)
fn.AddEdge(1, 2, 10)
fn.AddEdge(1, 3, 12)
fn.AddEdge(2, 1, 4)
fn.AddEdge(2, 4, 14)
fn.AddEdge(3, 2, 9)
fn.AddEdge(3, 5, 20)
fn.AddEdge(4, 3, 7)
fn.AddEdge(4, 5, 4)

maxFlow := fn.FindMaxFlow(0, 5)
fmt.Printf("Maximum flow: %d\n", maxFlow)
```

## Implementation Details

The package provides the following main components:

1. `FlowNetwork` struct: Represents a flow network with a residual graph and a flow graph.
2. `NewFlowNetwork(V int) *FlowNetwork`: Creates a new FlowNetwork with V vertices.
3. `AddEdge(u, v, c int)`: Adds a directed edge from u to v with capacity c.
4. `FindMaxFlow(source, sink int) int`: Finds the maximum flow from source to sink using the Ford-Fulkerson algorithm with Edmonds-Karp implementation.

The implementation uses adjacency matrices to represent the residual graph and the flow graph for simplicity and ease of understanding. For very large graphs, an adjacency list representation might be more memory-efficient.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple network
2. Complex network
3. Single path network
4. Disconnected network

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the maximum flow implementation under different network structures.

## Advantages and Limitations

Advantages:
- Correctly finds the maximum flow in any flow network
- Relatively simple to understand and implement
- Can be extended to solve other network flow problems

Limitations:
- May not be the most efficient for very large networks (there are more advanced algorithms with better time complexity)
- The adjacency matrix representation may consume too much memory for extremely large sparse graphs

## Applications

- Transportation and logistics (e.g., maximizing the flow of goods through a distribution network)
- Computer networking (e.g., maximizing data flow in a network)
- Bipartite matching problems
- Image segmentation in computer vision
- Sports tournament scheduling

Remember that while this implementation finds the maximum flow, it doesn't explicitly provide the flow assignment for each edge. If you need this information, you can extract it from the `Flow` matrix in the `FlowNetwork` struct after running the algorithm.

