package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/util"
	"github.com/jamestunnell/go-musicality/pkg/util/function"
)

func TestLinear(t *testing.T) {
	f := function.NewLinearFunctionFromPoints(pZeroZero, pOneOne)

	testFunctionAt(t, f, -0.5, -0.5)
	testFunctionAt(t, f, 0.0, 0.0)
	testFunctionAt(t, f, 0.5, 0.5)
	testFunctionAt(t, f, 1.0, 1.0)
	testFunctionAt(t, f, 1.5, 1.5)

	testFunctionSample(t, f, util.NewRange(0.0, 1.0), 0.5, []float64{0.0, 0.5, 1.0})
}

func TestLinear2(t *testing.T) {
	p0 := util.NewPoint(0.0, 0.5)
	p1 := util.NewPoint(1.0, 1.0)
	f := function.NewLinearFunctionFromPoints(p0, p1)

	testFunctionAt(t, f, -1, 0.0)
	testFunctionAt(t, f, -0.5, 0.25)
	testFunctionAt(t, f, 0.0, 0.5)
	testFunctionAt(t, f, 0.5, 0.75)
	testFunctionAt(t, f, 1.0, 1.0)
	testFunctionAt(t, f, 1.5, 1.25)
}
