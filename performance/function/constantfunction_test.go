package function_test

import (
	"math/big"
	"testing"

	"github.com/jamestunnell/go-musicality/performance/function"
)

func TestConstant(t *testing.T) {
	f := function.NewConstantFunction(2.5)

	zero := big.NewRat(0, 1)
	one := big.NewRat(1, 1)

	testFunctionAt(t, f, big.NewRat(-1, 1), 2.5)
	testFunctionAt(t, f, zero, 2.5)
	testFunctionAt(t, f, one, 2.5)

	testFunctionSample(t, f, function.NewRange(zero, one), half, []float64{2.5, 2.5, 2.5})
}
