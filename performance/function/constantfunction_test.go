package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/function"
)

func TestConstant(t *testing.T) {
	f := function.NewConstantFunction(2.5)

	testFunctionAt(t, f, rat.FromInt64(-1), 2.5)
	testFunctionAt(t, f, zero, 2.5)
	testFunctionAt(t, f, one, 2.5)

	testFunctionSample(t, f, function.NewRange(zero, one), half, []float64{2.5, 2.5, 2.5})
}
