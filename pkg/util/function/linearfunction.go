package function

import (
	"github.com/jamestunnell/go-musicality/pkg/util"
)

type LinearFunction struct {
	slope      float64
	yIntercept float64
}

func NewLinearFunction(slope, yIntercept float64) *LinearFunction {
	return &LinearFunction{slope: slope, yIntercept: yIntercept}
}

func NewLinearFunctionFromPoints(p0, p1 util.Point) *LinearFunction {
	slope := (p1.Y - p0.Y) / (p1.X - p0.X)
	intercept := p0.Y - slope*p0.X

	return NewLinearFunction(slope, intercept)
}

func (f *LinearFunction) At(x float64) float64 {
	return f.slope*x + f.yIntercept
}

func (f *LinearFunction) Domain() util.Range {
	return DomainAllFloat64
}
