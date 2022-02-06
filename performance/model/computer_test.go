package model_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func TestComputerNoChanges(t *testing.T) {
	startVal := 7.2
	c, err := model.NewComputer(startVal, change.Map{})

	assert.Nil(t, err)
	assert.NotNil(t, c)

	testComputerAt(t, c, function.DomainMin(), startVal)
	testComputerAt(t, c, big.NewRat(0, 1), startVal)
	testComputerAt(t, c, function.DomainMax(), startVal)
}

func TestComputerOneImmediateChange(t *testing.T) {
	offset := new(big.Rat).SetInt64(2)
	startVal := 20.0
	newVal := 10.0
	changes := change.Map{offset: change.NewImmediate(newVal)}
	c, err := model.NewComputer(startVal, changes)

	assert.Nil(t, err)
	assert.NotNil(t, c)

	testComputerAt(t, c, function.DomainMin(), startVal)
	testComputerAt(t, c, new(big.Rat).SetFloat64(1.99), startVal)
	testComputerAt(t, c, offset, newVal)
	testComputerAt(t, c, new(big.Rat).SetFloat64(2.01), newVal)
	testComputerAt(t, c, function.DomainMax(), newVal)
}

func TestComputerOneGradualChange(t *testing.T) {
	offset := new(big.Rat).SetInt64(5)
	startVal := 15.0
	newVal := 25.0
	dur := big.NewRat(10, 1)
	changes := change.Map{offset: change.New(newVal, dur)}
	c, err := model.NewComputer(startVal, changes)

	assert.Nil(t, err)
	assert.NotNil(t, c)

	testComputerAt(t, c, function.DomainMin(), startVal)
	testComputerAt(t, c, new(big.Rat).SetFloat64(4.99), startVal)
	testComputerAtNear(t, c, offset, startVal)
	testComputerAtNear(t, c, big.NewRat(10, 1), (startVal+newVal)/2.0)
	testComputerAtNear(t, c, big.NewRat(15, 1), newVal)
	testComputerAt(t, c, function.DomainMax(), newVal)
}

func testComputerAt(t *testing.T, c *model.Computer, x *big.Rat, expected float64) {
	y, err := function.At(c, x)

	assert.Nil(t, err)
	assert.Equal(t, expected, y)
}

func testComputerAtNear(t *testing.T, c *model.Computer, x *big.Rat, expected float64) {
	y, err := function.At(c, x)

	assert.Nil(t, err)
	assert.InDelta(t, expected, y, 1e-5)
}

// func testComputerOneGradualChange(t *testing.T, change *change.Change) {
// 	const (
// 		startVal = -10.0
// 		changeAt = 4
// 	)

// 	changeHalfDoneAt := changeAt + change.Duration/2
// 	changeDoneAt := changeAt + change.Duration
// 	changes := map[float64]*change.Change{changeAt: change}
// 	c, err := model.NewComputer(startVal, changes)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, c)

// 	testCases := map[float64]float64{
// 		(changeAt - 1e5):  startVal,
// 		(changeAt - 0.01): startVal,
// 		changeAt:          startVal,
// 		changeHalfDoneAt:  (startVal + change.EndValue) / 2.0,
// 		changeDoneAt:      change.EndValue,
// 		(changeAt + 1e5):  change.EndValue,
// 	}

// 	for x, y := range testCases {
// 		val, err := function.At(c, x)

// 		assert.Nil(t, err)
// 		assert.InDelta(t, y, val, 1e-5)
// 	}
// }
