package mazesolver

import (
    "errors"
)

type Point struct {
    X, Y int
}

type Maze [][]byte

func (m Maze) isValid(p Point) bool {
    return p.X >= 0 && p.X < len(m) && p.Y >= 0 && p.Y < len(m[0])
}

func (m Maze) findStart() (Point, error) {
    for i := range m {
        for j := range m[i] {
            if m[i][j] == 'S' {
                return Point{i, j}, nil
            }
        }
    }
    return Point{}, errors.New("start point not found")
}

func (m Maze) findEnd() (Point, error) {
    for i := range m {
        for j := range m[i] {
            if m[i][j] == 'E' {
                return Point{i, j}, nil
            }
        }
    }
    return Point{}, errors.New("end point not found")
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolveMazeDFS(maze Maze) ([]Point, error) {
    start, err := maze.findStart()
    if err != nil {
        return nil, err
    }
    end, err := maze.findEnd()
    if err != nil {
        return nil, err
    }

    visited := make([][]bool, len(maze))
    for i := range visited {
        visited[i] = make([]bool, len(maze[0]))
    }

    path := []Point{}
    if dfs(maze, start, end, visited, &path) {
        return path, nil
    }
    return nil, errors.New("no path found")
}

func dfs(maze Maze, current, end Point, visited [][]bool, path *[]Point) bool {
    if current == end {
        *path = append(*path, current)
        return true
    }

    if !maze.isValid(current) || visited[current.X][current.Y] || maze[current.X][current.Y] == '1' {
        return false
    }

    visited[current.X][current.Y] = true
    *path = append(*path, current)

    for _, dir := range directions {
        next := Point{current.X + dir.X, current.Y + dir.Y}
        if dfs(maze, next, end, visited, path) {
            return true
        }
    }

    *path = (*path)[:len(*path)-1]
    return false
}

func SolveMazeBFS(maze Maze) ([]Point, error) {
    start, err := maze.findStart()
    if err != nil {
        return nil, err
    }
    end, err := maze.findEnd()
    if err != nil {
        return nil, err
    }

    visited := make([][]bool, len(maze))
    for i := range visited {
        visited[i] = make([]bool, len(maze[0]))
    }

    queue := []Point{start}
    parent := make(map[Point]Point)

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == end {
            return reconstructPath(parent, end), nil
        }

        for _, dir := range directions {
            next := Point{current.X + dir.X, current.Y + dir.Y}
            if maze.isValid(next) && !visited[next.X][next.Y] && maze[next.X][next.Y] != '1' {
                visited[next.X][next.Y] = true
                queue = append(queue, next)
                parent[next] = current
            }
        }
    }

    return nil, errors.New("no path found")
}

func reconstructPath(parent map[Point]Point, end Point) []Point {
    path := []Point{end}
    for p, ok := parent[end]; ok; p, ok = parent[p] {
        path = append([]Point{p}, path...)
    }
    return path
}
