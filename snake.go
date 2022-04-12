package main

const (
	Up    int = 0
	Down      = 1
	Left      = 2
	Right     = 3
)

type Snake struct {
	Body      []Point
	Direction int
	Head      Point
}

type Point struct {
	X int
	Y int
}

func (s *Snake) Move() {
	for _, b := range s.Body {
		b.X++
	}
}

func (s *Snake) Render() []Point {
	return s.Body
}
