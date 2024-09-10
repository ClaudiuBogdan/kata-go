# A* Search Algorithm for Pathfinding

## Problem Description

Implement the A* search algorithm for pathfinding in a 2D grid. The algorithm should find the shortest path between a start node and a goal node, considering obstacles and using a heuristic function to guide the search.

## Approach

The A* algorithm is a best-first search algorithm that combines the benefits of Dijkstra's algorithm (which always finds the shortest path) and greedy best-first search (which is fast but doesn't always find the shortest path). It uses a heuristic function to estimate the cost from any given node to the goal, which helps guide the search towards the goal more efficiently.

The key steps of the A* algorithm are:

1. Initialize the open list with the start node.
2. While the open list is not empty:
   a. Get the node with the lowest F score from the open list.
   b. If this node is the goal, reconstruct and return the path.
   c. Move this node to the closed list.
   d. For each neighbor of the current node:
      - If it's in the closed list, skip it.
      - If it's not in the open list, add it.
      - If it's already in the open list, check if this path is better (lower G score).
3. If the open list is empty and the goal hasn't been reached, there's no path.

## Complexity Analysis

- Time Complexity: O(b^d), where b is the branching factor (average number of successors per state) and d is the depth of the shortest path.
- Space Complexity: O(b^d), as we need to store all generated nodes.

In practice, the performance of A* can vary greatly depending on the heuristic function used. A good heuristic can significantly reduce the number of nodes explored.

## Usage

```go
grid := NewGrid(10, 10)
// Set up obstacles
grid.Nodes[2][2].Walkable = false
grid.Nodes[2][3].Walkable = false
grid.Nodes[2][4].Walkable = false

start := grid.Nodes[0][0]
goal := grid.Nodes[9][9]

path := AStar(grid, start, goal)

if path != nil {
    fmt.Println("Path found:")
    for _, node := range path {
        fmt.Printf("(%d, %d) ", node.X, node.Y)
    }
} else {
    fmt.Println("No path found")
}
```

## Implementation Details

The package provides the following main components:

1. `Node`: Represents a point in the grid with properties like position, G score (cost from start), H score (heuristic estimate to goal), F score (G + H), parent node, and walkability.
2. `Grid`: Represents the search space as a 2D grid of nodes.
3. `PriorityQueue`: A min-heap implementation for efficiently selecting the node with the lowest F score.
4. `AStar`: The main function that performs the A* search algorithm.
5. `Heuristic`: Calculates the estimated cost from a node to the goal using Euclidean distance.

The implementation uses a priority queue to efficiently select the node with the lowest F score, which is crucial for the performance of the A* algorithm.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Simple path finding
2. Path finding with obstacles
3. Scenarios where no path exists
4. Heuristic function accuracy

To run the tests, use the following command:

```bash
go test
```

## Advantages of A* Search

- Generally faster than Dijkstra's algorithm for pathfinding problems.
- Guaranteed to find the shortest path if one exists (assuming an admissible heuristic).
- Can handle large search spaces efficiently with a good heuristic.

## Applications

- Video game pathfinding
- Robotics and motion planning
- Geographic Information Systems (GIS) for route planning
- Network routing protocols

The A* algorithm is versatile and can be adapted to various pathfinding problems beyond simple 2D grids. By modifying the heuristic function and the way neighbors are generated, it can be applied to different types of graphs and search spaces.
