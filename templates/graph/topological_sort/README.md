# Topological Sorting of a Directed Acyclic Graph (DAG)

## Problem Description

Implement an algorithm to perform topological sorting on a directed acyclic graph (DAG). A topological sort of a DAG is a linear ordering of its vertices such that for every directed edge (u, v), vertex u comes before vertex v in the ordering. Topological sorting is only possible for DAGs.

## Approach

This implementation uses a depth-first search (DFS) approach to perform topological sorting. The algorithm works as follows:

1. Create a recursive function that takes the current vertex, visited array, and a stack to store the result.
2. Mark the current node as visited.
3. Recursively call the function for all the vertices adjacent to this vertex.
4. Push the vertex to a stack after calling the recursive function for all of its adjacent vertices.
5. After the recursive calls are complete, pop the stack to get the topological order.

The implementation also includes cycle detection to ensure the graph is acyclic.

## Complexity Analysis

- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges.
- Space Complexity: O(V) for the visited array, recursion stack, and output stack.

## Usage

```go
g := NewGraph(4)
g.AddEdge(0, 1)
g.AddEdge(0, 2)
g.AddEdge(1, 3)
g.AddEdge(2, 3)

order, success := g.TopologicalSort()
if success {
    fmt.Println("Topological Order:", order)
} else {
    fmt.Println("The graph contains a cycle and cannot be topologically sorted.")
}
```

## Implementation Details

The package provides the following main components:

1. `Graph` struct: Represents a directed graph using an adjacency list.
2. `NewGraph(V int) *Graph`: Creates a new Graph with V vertices.
3. `AddEdge(u, v int)`: Adds a directed edge from u to v.
4. `TopologicalSort() ([]int, bool)`: Performs topological sorting on the graph, returning the sorted order and a boolean indicating success (false if a cycle is detected).

The implementation uses depth-first search (DFS) to efficiently perform topological sorting and detect cycles.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple DAG
2. Complex DAG
3. Single vertex graph
4. Graph with a cycle

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the topological sorting implementation under different graph structures.

## Advantages and Limitations

Advantages:
- Efficiently performs topological sorting on directed acyclic graphs
- Detects cycles in the graph, ensuring the result is valid
- Linear time complexity O(V + E)

Limitations:
- Only works for directed acyclic graphs
- The order of vertices with the same precedence may vary between runs

## Applications

- Dependency resolution in build systems or package managers
- Task scheduling in project management
- Course prerequisite planning in academic systems
- Data processing pipelines
- Static code analysis for detecting circular dependencies

Remember that while this implementation uses DFS, there are other approaches to topological sorting, such as Kahn's algorithm, which uses a queue-based approach. The choice of algorithm may depend on specific requirements or constraints of your application.

## Video

- [Topological Sort Algorithm | Graph Theory](https://www.youtube.com/watch?v=eL-KzMXSXXI&t=18s)

- [Topological Sort Visualized and Explained](https://www.youtube.com/watch?v=7J3GadLzydI)

