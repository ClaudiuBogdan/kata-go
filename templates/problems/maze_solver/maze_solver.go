package mazesolver

type Point struct {
    X, Y int
}

type Maze [][]byte

func (m Maze) isValid(p Point) bool {

}

func (m Maze) findStart() (Point, error) {

}

func (m Maze) findEnd() (Point, error) {

}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func SolveMazeDFS(maze Maze) ([]Point, error) {

}

func dfs(maze Maze, current, end Point, visited [][]bool, path *[]Point) bool {

}

func SolveMazeBFS(maze Maze) ([]Point, error) {

}

func reconstructPath(parent map[Point]Point, end Point) []Point {

}
