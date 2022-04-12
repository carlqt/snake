package main

func main() {
	stage := NewStage(10, 20)
	snake := Snake{
		Body: []Point{
			{Y: 4, X: 1},
			{Y: 4, X: 2},
			{Y: 4, X: 3},
			{Y: 4, X: 4},
			{Y: 4, X: 5},
			{Y: 4, X: 6},
			{Y: 4, X: 7},
			{Y: 4, X: 8},
		},
	}

	stage.Snake = snake
	stage.Render()
}
