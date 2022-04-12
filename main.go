package main

import (
	"time"
)

func main() {
	stage := NewStage(10, 20)
	snake := Snake{
		Body: []Point{
			{Y: 4, X: 1},
			{Y: 4, X: 2},
			{Y: 4, X: 3},
		},
	}

	stage.Clear()
	stage.Snake = snake

	for x := 5; x >= 0; x-- {
		stage.Render()
		stage.Snake.Move()
		// fmt.Print("\033[u\033[K")
		time.Sleep(500 * time.Millisecond)
	}
}
