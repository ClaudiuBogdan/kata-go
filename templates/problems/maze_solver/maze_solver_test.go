package mazesolver

import (
	"reflect"
	"testing"
)

func TestMazeSolver(t *testing.T) {
    testCases := []struct {
        name     string
        maze     Maze
        expected []Point
    }{
        {
            name: "Simple maze",
            maze: Maze{
                {'S', '0', '1', '1'},
                {'1', '0', '0', '1'},
                {'1', '0', '1', '1'},
                {'1', '0', 'E', '1'},
            },
            expected: []Point{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {3, 1}, {3, 2}},
        },
        {
            name: "Maze with no solution",
            maze: Maze{
                {'S', '1', '1', '1'},
                {'1', '1', '0', '1'},
                {'1', '0', '1', '1'},
                {'1', '0', 'E', '1'},
            },
            expected: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name+" (DFS)", func(t *testing.T) {
            path, err := SolveMazeDFS(tc.maze)
            if tc.expected == nil {
                if err == nil {
                    t.Errorf("Expected error, but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if !reflect.DeepEqual(path, tc.expected) {
                    t.Errorf("Expected path %v, but got %v", tc.expected, path)
                }
            }
        })

        t.Run(tc.name+" (BFS)", func(t *testing.T) {
            path, err := SolveMazeBFS(tc.maze)
            if tc.expected == nil {
                if err == nil {
                    t.Errorf("Expected error, but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if !reflect.DeepEqual(path, tc.expected) {
                    t.Errorf("Expected path %v, but got %v", tc.expected, path)
                }
            }
        })
    }
}

func TestInvalidMaze(t *testing.T) {
    testCases := []struct {
        name string
        maze Maze
    }{
        {
            name: "Maze without start",
            maze: Maze{
                {'0', '0', '1', '1'},
                {'1', '0', '0', '1'},
                {'1', '0', '1', '1'},
                {'1', '0', 'E', '1'},
            },
        },
        {
            name: "Maze without end",
            maze: Maze{
                {'S', '0', '1', '1'},
                {'1', '0', '0', '1'},
                {'1', '0', '1', '1'},
                {'1', '0', '0', '1'},
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name+" (DFS)", func(t *testing.T) {
            _, err := SolveMazeDFS(tc.maze)
            if err == nil {
                t.Errorf("Expected error, but got none")
            }
        })

        t.Run(tc.name+" (BFS)", func(t *testing.T) {
            _, err := SolveMazeBFS(tc.maze)
            if err == nil {
                t.Errorf("Expected error, but got none")
            }
        })
    }
}
