package mazesolver

import "fmt"

type Point struct {
	X, Y int
}

type Maze [][]byte

func (m Maze) isValid(p Point) bool {
	if p.X < 0 || p.X >= len(m) {
		return false
	}

	if p.Y < 0 || p.Y >= len(m[0]) {
		return false
	}

	if m[p.X][p.Y] == '1' {
		return false
	}

	return true
}

func (m Maze) findStart() (Point, error) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 'S' {
				return Point{X: i, Y: j}, nil
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
				return Point{X: i, Y: j}, nil
			}
		}
	}
	var zero Point
	return zero, fmt.Errorf("end not found")
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolveMazeDFS(maze Maze) ([]Point, error) {
	visited := make([][]bool, len(maze))
	path := make([]Point, 0)

	var zero []Point
	start, err := maze.findStart()
	if err != nil {
		return zero, err
	}
	end, err := maze.findEnd()
	if err != nil {
		return zero, err
	}

	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	if dfs(maze, start, end, visited, &path) {
        // Revert path
        newPath := make([]Point, len(path))
        for i := 0; i < len(path); i++{
            newPath[i] = path[len(path) - i - 1]
        }
		return newPath, nil
	}
	return zero, fmt.Errorf("path not found")
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
	visited := make([][]bool, len(maze))
	path := make(map[Point]Point, 0)

	var zero []Point
	start, err := maze.findStart()
	if err != nil {
		return zero, err
	}
	end, err := maze.findEnd()
	if err != nil {
		return zero, err
	}

	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

    queue := []Point{start}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        visited[current.X][current.Y] = true

        if current == end {
            return reconstructPath(path, end), nil
        }

        for _, dir := range directions {
            next := Point{
                current.X + dir.X,
                current.Y + dir.Y,
            }

            if maze.isValid(next) && !visited[next.X][next.Y]{
                queue = append(queue, next)
                path[next] = current
            }
        }
    }

    
    return zero, fmt.Errorf("path not found")
}

func reconstructPath(parent map[Point]Point, end Point) []Point {
    path := make([]Point, 0)
    for currPoint, hasParent := end, true;  hasParent; currPoint, hasParent = parent[currPoint]{
        path = append([]Point{currPoint}, path...)
    }
    return path
}
