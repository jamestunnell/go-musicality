package function

import (
	"github.com/jamestunnell/go-musicality/notation/rat"
)

type Point struct {
	X rat.Rat
	Y float64
}

func NewPoint(x rat.Rat, y float64) Point {
	return Point{X: x, Y: y}
}
