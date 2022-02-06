package function

import "math/big"

type LinearFunction struct {
	domain     Range
	slope      float64
	yIntercept float64
}

func NewLinearFunction(slope, yIntercept float64) *LinearFunction {
	return &LinearFunction{
		domain:     DomainAll(),
		slope:      slope,
		yIntercept: yIntercept,
	}
}

func NewLinearFunctionFromPoints(p0, p1 Point) *LinearFunction {
	xDelta, _ := new(big.Rat).Sub(p1.X, p0.X).Float64()
	slope := (p1.Y - p0.Y) / xDelta
	x0, _ := p0.X.Float64()
	intercept := p0.Y - slope*x0

	return NewLinearFunction(slope, intercept)
}

func (f *LinearFunction) At(x *big.Rat) float64 {
	xFlt, _ := x.Float64()

	return f.slope*xFlt + f.yIntercept
}

func (f *LinearFunction) Domain() Range {
	return f.domain
}
