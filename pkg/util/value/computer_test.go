package value_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/pkg/util/value"

	"github.com/stretchr/testify/assert"
)

func TestComputerNoChanges(t *testing.T) {
	c, err := value.NewComputer(7.2, map[float64]*value.Change{})

	assert.Nil(t, err)
	assert.NotNil(t, c)

	testCases := map[float64]float64{-1e75: 7.2, 0: 7.2, 1e75: 7.2}

	for x, y := range testCases {
		val, err := function.At(c, x)

		assert.Nil(t, err)
		assert.Equal(t, y, val)
	}
}

func TestComputerOneGradualChange(t *testing.T) {
	for _, endVal := range []float64{2.7, -1546, 55.96} {
		for _, dur := range []float64{0.01, 1, 7.44, 1456.2} {
			change, err := value.NewLinearChange(endVal, dur)

			assert.NotNil(t, change)
			assert.Nil(t, err)

			testComputerOneGradualChange(t, change)

			change, err = value.NewSigmoidChange(endVal, dur)

			assert.NotNil(t, change)
			assert.Nil(t, err)

			testComputerOneGradualChange(t, change)
		}
	}
}

func testComputerOneGradualChange(t *testing.T, change *value.Change) {
	const (
		startVal = -10.0
		changeAt = 4
	)

	changeHalfDoneAt := changeAt + change.Duration/2
	changeDoneAt := changeAt + change.Duration
	changes := map[float64]*value.Change{changeAt: change}
	c, err := value.NewComputer(startVal, changes)

	assert.Nil(t, err)
	assert.NotNil(t, c)

	testCases := map[float64]float64{
		(changeAt - 1e5):  startVal,
		(changeAt - 0.01): startVal,
		changeAt:          startVal,
		changeHalfDoneAt:  (startVal + change.EndValue) / 2.0,
		changeDoneAt:      change.EndValue,
		(changeAt + 1e5):  change.EndValue,
	}

	for x, y := range testCases {
		val, err := function.At(c, x)

		assert.Nil(t, err)
		assert.InDelta(t, y, val, 1e-5)
	}
}
