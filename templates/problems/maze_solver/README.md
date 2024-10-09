# Maze Solver

## Problem Description

Given a 2D maze represented as a grid of characters, find a path from the start point 'S' to the end point 'E'. The maze consists of open paths ('0'), walls ('1'), a start point ('S'), and an end point ('E'). The goal is to find a valid path from 'S' to 'E', if one exists.

## Approach: Depth-First Search (DFS) and Breadth-First Search (BFS)

This package implements two popular graph traversal algorithms to solve the maze:

1. Depth-First Search (DFS): Explores as far as possible along each branch before backtracking.
2. Breadth-First Search (BFS): Explores all the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level.

## Complexity Analysis

- Time Complexity: O(N*M) for both DFS and BFS, where N and M are the dimensions of the maze.
- Space Complexity: O(N*M) for both algorithms to store the visited cells and the path.

## Implementation Details

The package provides two main functions:

`SolveMazeDFS(maze Maze) ([]Point, error)`: Solves the maze using Depth-First Search.
`SolveMazeBFS(maze Maze) ([]Point, error)`: Solves the maze using Breadth-First Search.

Both functions take a `Maze` type (which is a 2D byte slice) as input and return a slice of `Point` structures representing the path from start to end, or an error if no path is found.

The `Point` structure represents a coordinate in the maze:

```go
type Point struct {
    X, Y int
}
```

## Usage

```go
maze := Maze{
    {'S', '0', '1', '1'},
    {'1', '0', '0', '1'},
    {'1', '0', '1', '1'},
    {'1', '0', '0', 'E'},
}

path, err := SolveMazeDFS(maze)
if err != nil {
    fmt.Println("No path found")
} else {
    fmt.Println("Path found:", path)
}
```

## Testing

The implementation should include a comprehensive test suite covering various scenarios:

1. Maze with a valid path
2. Maze with no valid path
3. Large maze
4. Edge cases (1x1 maze, maze with only start and end points, etc.)

To run the tests, use the following command:

```bash
go test
```

## Advantages of DFS and BFS for Maze Solving

1. Guaranteed to find a path if one exists
2. DFS can be more memory-efficient for deep mazes
3. BFS always finds the shortest path
4. Both algorithms are relatively simple to implement

## Applications

- Pathfinding in video games
- Robot navigation
- Circuit board design
- Network routing

## Visual Representation

Here's a visual representation of a simple maze and its solution:

```
S 0 1 1    S * 1 1
1 0 0 1    1 * 0 1
1 0 1 1 -> 1 * 1 1
1 0 0 E    1 0 0 E
```

The '*' characters represent the found path.

## Variations and Extensions

1. Finding the shortest path (guaranteed with BFS)
2. Solving mazes with multiple end points
3. Solving 3D mazes
4. Adding weighted paths (turning it into a weighted graph problem)

## Implementation Notes

- The implementation assumes that the input maze is valid (contains 'S' and 'E').
- Error handling is included for cases where start or end points are not found.
- The solution provides the path as a series of coordinates.

## Limitations

- Does not handle mazes with cycles efficiently (though this is rarely an issue in most maze problems)
- May not be the most efficient for extremely large mazes

## Related Problems

- Shortest Path in a Graph
- Island Count
- Word Search
- Robot Room Cleaner
- Solving Sudoku

Remember that while DFS and BFS are excellent for many maze-solving scenarios, other algorithms like A* or Dijkstra's algorithm might be more appropriate for weighted mazes or when finding the optimal path is crucial.
