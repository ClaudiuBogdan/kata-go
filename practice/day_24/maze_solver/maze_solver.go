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
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 'S' {
				return Point{X: i, Y: j}, nil
			}
		}
	}
	return Point{}, fmt.Errorf("start not founds")
}

func (m Maze) findEnd() (Point, error) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 'E' {
				return Point{X: i, Y: j}, nil
			}
		}
	}
	return Point{}, fmt.Errorf("end not found")
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolveMazeDFS(maze Maze) ([]Point, error) {
	result := make([]Point, 0)
	start, err := maze.findStart()
	if err != nil {
		return result, err
	}
	end, err := maze.findEnd()
	if err != nil {
		return result, err
	}
	visited := make([][]bool, len(maze))

	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}

    visited[start.X][start.Y] = true

	if dfs(maze, start, end, visited, &result) {
		reversed := make([]Point, len(result))
		for i := range result {
			reversed[i] = result[len(result)-i-1]
		}
		return reversed, nil
	}
	return []Point{}, fmt.Errorf("path not found")
}

func dfs(maze Maze, current, end Point, visited [][]bool, path *[]Point) bool {
	if current == end {
		*path = append(*path, end)
		return true
	}

	for _, dir := range directions {
		next := Point{
			X: current.X + dir.X,
			Y: current.Y + dir.Y,
		}

		if maze.isValid(next) && !visited[next.X][next.Y] {
			visited[next.X][next.Y] = true

			if dfs(maze, next, end, visited, path) {
                // TODO: add current node
				*path = append(*path, current)
				return true
			}
		}
	}

	return false
}
      
func SolveMazeBFS(maze Maze) ([]Point, error) {
	parent := make(map[Point]Point)

	start, err := maze.findStart()
	if err != nil {
		return []Point{}, fmt.Errorf("path not found")
	}

	end, err := maze.findEnd()
	if err != nil {
		return []Point{}, fmt.Errorf("path not found")
	}


	visited := make([][]bool, len(maze))

    // TODO: initialize visited matrix.
    for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}

	queue := []Point{start}
	visited[start.X][start.Y] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return reconstructPath(parent, end), nil
		}

		for _, dir := range directions {
			next := Point{
				X: current.X + dir.X,
				Y: current.Y + dir.Y,
			}

			if maze.isValid(next) && !visited[next.X][next.Y] {
				queue = append(queue, next)
				visited[next.X][next.Y] = true
				parent[next] = current
			}
		}
	}
	return nil, fmt.Errorf("path not found")
}

func reconstructPath(parent map[Point]Point, end Point) []Point {
    path := []Point{end}

    for current, ok := parent[end]; ok; current, ok = parent[current]{
        path = append([]Point{current}, path...)
    }
    return path
}
