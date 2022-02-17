package function

import "github.com/jamestunnell/go-musicality/common/rat"

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
	xDelta := p1.X.Sub(p0.X).Float64()
	slope := (p1.Y - p0.Y) / xDelta
	x0 := p0.X.Float64()
	intercept := p0.Y - slope*x0

	return NewLinearFunction(slope, intercept)
}

func (f *LinearFunction) At(x rat.Rat) float64 {
	return f.slope*x.Float64() + f.yIntercept
}

func (f *LinearFunction) Domain() Range {
	return f.domain
}
