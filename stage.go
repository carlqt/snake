package main

import "fmt"

// 10x20
// type Stage [][]string

type Stage struct {
	Grid [][]string
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

func (s *Stage) DisplayObject(obj []Point) {
	for _, p := range obj {
		x := p.X
		y := p.Y

		s.Grid[y][x] = "█"
		// (*s.Grid)[y][x] = "█"
	}
}

func (stage *Stage) Render() {
	obj := []Point{
		{Y: 4, X: 1},
		{Y: 4, X: 2},
		{Y: 4, X: 3},
		{Y: 4, X: 4},
		{Y: 4, X: 5},
		{Y: 4, X: 6},
		{Y: 4, X: 7},
		{Y: 4, X: 8},
	}

	stage.DisplayObject(obj)

	for _, row := range stage.Grid {
		for _, column := range row {
			fmt.Print(column)
		}

		fmt.Println("")
	}
}
