package function_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/util"
	"github.com/jamestunnell/go-musicality/pkg/util/function"
	"github.com/stretchr/testify/assert"
)

var (
	pZeroZero = util.NewPoint(0.0, 0.0)
	pOneOne   = util.NewPoint(1.0, 1.0)
)

func TestSigmoid(t *testing.T) {
	f := function.NewSigmoidFunction(pZeroZero, pOneOne)

	testFunctionAt(t, f, 0.0, 0.0)
	testFunctionAt(t, f, 0.5, 0.5)
	testFunctionAt(t, f, 1.0, 1.0)

	testFunctionSample(t, f, util.NewRange(0.0, 1.0), 0.5, []float64{0.0, 0.5, 1.0})
}

func TestSigmoidOutOfDomain(t *testing.T) {
	f := function.NewSigmoidFunction(pZeroZero, pOneOne)

	_, err := function.At(f, -0.01)

	assert.NotNil(t, err)

	_, err = function.At(f, 1.01)

	assert.NotNil(t, err)
}
