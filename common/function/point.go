package function

import "math/big"

type Point struct {
	X *big.Rat
	Y float64
}

func NewPoint(x *big.Rat, y float64) Point {
	return Point{X: x, Y: y}
}
