package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/util"
	"github.com/jamestunnell/go-musicality/pkg/util/function"
)

func TestConstant(t *testing.T) {
	f := function.NewConstantFunction(2.5)

	testFunctionAt(t, f, -1e6, 2.5)
	testFunctionAt(t, f, 0.0, 2.5)
	testFunctionAt(t, f, 1e6, 2.5)

	testFunctionSample(t, f, util.NewRange(0.0, 1.0), 0.5, []float64{2.5, 2.5, 2.5})
}
