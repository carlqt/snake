package main

import "fmt"

// 10x20
// type Stage [][]string
type Renderer interface {
	Render() []Point
}

type Stage struct {
	Grid    [][]string
	Objects []Renderer
	Snake   Snake
}

func NewStage(l int, w int) *Stage {
	// stage := make(Stage, l)
	var stage Stage

	stage.Grid = make([][]string, l)

	for i := range stage.Grid {
		stage.Grid[i] = make([]string, w)
	}

	return &stage
}

func (s *Stage) SetSnake() {
	for _, p := range s.Snake.Body {
		x := p.X
		y := p.Y

		s.Grid[y][x] = "█"
	}
}

func (stage *Stage) Render() {
	stage.SetSnake()

	for _, row := range stage.Grid {
		for _, column := range row {
			fmt.Print(column)
		}

		fmt.Println("")
	}
}
