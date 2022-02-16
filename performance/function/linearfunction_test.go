package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
)

var (
	negHalf = rat.FromFloat64(-0.5)
)

func TestLinear(t *testing.T) {
	p0 := function.NewPoint(zero, 0.0)
	p1 := function.NewPoint(one, 1.0)
	f := function.NewLinearFunctionFromPoints(p0, p1)

	testFunctionAt(t, f, negHalf, -0.5)
	testFunctionAt(t, f, zero, 0.0)
	testFunctionAt(t, f, half, 0.5)
	testFunctionAt(t, f, one, 1.0)

	testFunctionSample(t, f, function.NewRange(zero, one), half, []float64{0.0, 0.5, 1.0})
}

func TestLinear2(t *testing.T) {
	p0 := function.NewPoint(zero, 0.5)
	p1 := function.NewPoint(one, 1.0)
	f := function.NewLinearFunctionFromPoints(p0, p1)

	testFunctionAt(t, f, negOne, 0.0)
	testFunctionAt(t, f, negHalf, 0.25)
	testFunctionAt(t, f, zero, 0.5)
	testFunctionAt(t, f, half, 0.75)
	testFunctionAt(t, f, one, 1.0)
}
