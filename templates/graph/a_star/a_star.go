package astar

import (
	"container/heap"
	"math"
)

// Node represents a point in the grid
type Node struct {
	X, Y     int
	G, H, F  float64
	Parent   *Node
	Walkable bool
}

// Grid represents the search space
type Grid struct {
	Nodes [][]*Node
	Width, Height int
}

// PriorityQueue implements heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].F < pq[j].F }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// NewGrid creates a new Grid with the given width and height
func NewGrid(width, height int) *Grid {
	grid := &Grid{
		Width:  width,
		Height: height,
		Nodes:  make([][]*Node, height),
	}
	for y := 0; y < height; y++ {
		grid.Nodes[y] = make([]*Node, width)
		for x := 0; x < width; x++ {
			grid.Nodes[y][x] = &Node{X: x, Y: y, Walkable: true}
		}
	}
	return grid
}

// GetNeighbors returns the neighboring nodes of the given node
func (g *Grid) GetNeighbors(node *Node) []*Node {
	neighbors := make([]*Node, 0, 8)
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			newX, newY := node.X+x, node.Y+y
			if newX >= 0 && newX < g.Width && newY >= 0 && newY < g.Height {
				neighbor := g.Nodes[newY][newX]
				if neighbor.Walkable {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}
	return neighbors
}

// Heuristic calculates the estimated cost from node to goal
func Heuristic(node, goal *Node) float64 {
	dx := float64(abs(node.X - goal.X))
	dy := float64(abs(node.Y - goal.Y))
	return math.Sqrt(dx*dx + dy*dy)
}

// AStar performs A* pathfinding on the given grid
func AStar(grid *Grid, start, goal *Node) []*Node {
	openList := &PriorityQueue{}
	heap.Init(openList)
	heap.Push(openList, start)

	closedSet := make(map[*Node]bool)
	start.G = 0
	start.H = Heuristic(start, goal)
	start.F = start.G + start.H

	for openList.Len() > 0 {
		current := heap.Pop(openList).(*Node)
		if current == goal {
			return reconstructPath(current)
		}

		closedSet[current] = true

		for _, neighbor := range grid.GetNeighbors(current) {
			if closedSet[neighbor] {
				continue
			}

			tentativeG := current.G + Heuristic(current, neighbor)

			if !contains(openList, neighbor) {
				heap.Push(openList, neighbor)
			} else if tentativeG >= neighbor.G {
				continue
			}

			neighbor.Parent = current
			neighbor.G = tentativeG
			neighbor.H = Heuristic(neighbor, goal)
			neighbor.F = neighbor.G + neighbor.H
		}
	}

	return nil // No path found
}

func reconstructPath(node *Node) []*Node {
	path := []*Node{node}
	for node.Parent != nil {
		node = node.Parent
		path = append([]*Node{node}, path...)
	}
	return path
}

func contains(pq *PriorityQueue, node *Node) bool {
	for _, n := range *pq {
		if n == node {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
