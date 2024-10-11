package mazesolver

import "fmt"

type Point struct {
	X, Y int
}

type Maze [][]byte

func (m Maze) isValid(p Point) bool {
	return p.X >= 0 && p.X < len(m) && p.Y >= 0 && p.Y < len(m[0]) && m[p.X][p.Y] != '1'
}

func (m Maze) findStart() (Point, error) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'S' {
				return Point{
					X: i,
					Y: j,
				}, nil
			}
		}
	}
	var zero Point
	return zero, fmt.Errorf("start not found")
}

func (m Maze) findEnd() (Point, error) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'E' {
				return Point{
					X: i,
					Y: j,
				}, nil
			}
		}
	}
	var zero Point
	return zero, fmt.Errorf("end not found")
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolveMazeDFS(maze Maze) ([]Point, error) {
	path := make([]Point, 0)
	start, err := maze.findStart()
	if err != nil {
		return nil, err
	}
	end, err := maze.findEnd()
	if err != nil {
		return nil, err
	}

	visited := make([][]bool, len(maze))

	for i := range len(visited) {
		visited[i] = make([]bool, len(maze[i]))
	}

	if dfs(maze, start, end, visited, &path) {
		// Revert path
		newPath := make([]Point, len(path))
		for i := 0; i < len(path); i++ {
			newPath[i] = path[len(path)-i-1]
		}
		return newPath, nil
	}
	return nil, fmt.Errorf("path not found")
}

func dfs(maze Maze, current, end Point, visited [][]bool, path *[]Point) bool {
	visited[current.X][current.Y] = true

	if current == end {
		*path = append(*path, current)
		return true
	}

	for _, dir := range directions {
		next := Point{
			X: current.X + dir.X,
			Y: current.Y + dir.Y,
		}

		if maze.isValid(next) && !visited[next.X][next.Y] && dfs(maze, next, end, visited, path) {
			*path = append(*path, current)
			return true
		}
	}
	return false
}

func SolveMazeBFS(maze Maze) ([]Point, error) {
	parent := make(map[Point]Point)

	start, err := maze.findStart()
	if err != nil {
		return nil, err
	}
	end, err := maze.findEnd()
	if err != nil {
		return nil, err
	}

	visited := make([][]bool, len(maze))

	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(maze[i]))
	}

	queue := []Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return reconstructPath(parent, end), nil
		}

		visited[current.X][current.Y] = true

		for _, dir := range directions {
			next := Point{
				X: current.X + dir.X,
				Y: current.Y + dir.Y,
			}

			if maze.isValid(next) && !visited[next.X][next.Y] {
				queue = append(queue, next)
				parent[next] = current
			}
		}
	}

	return nil, fmt.Errorf("path not found")
}

func reconstructPath(parent map[Point]Point, end Point) []Point {
	path := []Point{end}

	for current, ok := parent[end]; ok; current, ok = parent[current] {
		path = append([]Point{current}, path...)
	}
	return path
}
