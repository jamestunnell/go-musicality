package rat_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type X struct {
	Y *big.Rat
}

func TestBigRat(t *testing.T) {
	x := &X{Y: big.NewRat(1, 2)}

	d, err := json.Marshal(x)

	require.NoError(t, err)

	var x2 X

	require.NoError(t, json.Unmarshal(d, &x2))

	t.Logf("x1.Y: %s", x.Y.String())
	t.Logf("x2.Y: %s", x2.Y.String())
}

func TestFromFloat64(t *testing.T) {
	r := rat.FromFloat64(2.5)

	f, _ := r.Float64()

	assert.InDelta(t, 2.5, f, 1e-10)
}

func TestFromInt64(t *testing.T) {
	r := big.NewRat(73, 1)
	r2 := rat.FromInt64(73)

	assert.True(t, rat.IsEqual(r2, r))
}

func TestFromUint64(t *testing.T) {
	r := big.NewRat(73, 1)
	r2 := rat.FromUint64(73)

	assert.True(t, rat.IsEqual(r2, r))
}

func TestPositiveZeroNegative(t *testing.T) {
	testPositiveZeroNegative(t, big.NewRat(1, 2), true, false, false)
	testPositiveZeroNegative(t, rat.Zero(), false, true, false)
	testPositiveZeroNegative(t, big.NewRat(-1, 2), false, false, true)
}

func TestCompares(t *testing.T) {
	r2 := big.NewRat(2, 7)
	r1 := rat.Sub(r2, big.NewRat(1, 100))
	r3 := rat.Add(r2, big.NewRat(1, 100))

	assert.True(t, rat.IsLess(r1, r2))
	assert.True(t, rat.IsLessEqual(r1, r2))
	assert.True(t, rat.IsLessEqual(r1, r1))
	assert.True(t, rat.IsEqual(r1, r1))
	assert.True(t, rat.IsGreaterEqual(r1, r1))
	assert.True(t, rat.IsGreaterEqual(r2, r1))
	assert.True(t, rat.IsGreater(r2, r1))

	assert.True(t, rat.IsLess(r2, r3))
	assert.True(t, rat.IsLessEqual(r2, r3))
	assert.True(t, rat.IsLessEqual(r2, r2))
	assert.True(t, rat.IsEqual(r2, r2))
	assert.True(t, rat.IsGreaterEqual(r2, r2))
	assert.True(t, rat.IsGreaterEqual(r3, r2))
	assert.True(t, rat.IsGreater(r3, r2))
}

func TestMaths(t *testing.T) {
	r1 := big.NewRat(1, 2)
	r2 := big.NewRat(3, 4)

	assert.True(t, rat.IsEqual(rat.Add(r1, r2), big.NewRat(5, 4)))
	assert.True(t, rat.IsEqual(rat.Sub(r1, r2), big.NewRat(-1, 4)))
	assert.True(t, rat.IsEqual(rat.Mul(r1, r2), big.NewRat(3, 8)))
	assert.True(t, rat.IsEqual(rat.Div(r1, r2), big.NewRat(2, 3)))
}

func testPositiveZeroNegative(
	t *testing.T,
	r *big.Rat,
	pos, zero, neg bool) {
	assert.Equal(t, pos, rat.IsPositive(r))
	assert.Equal(t, zero, rat.IsZero(r))
	assert.Equal(t, neg, rat.IsNegative(r))
}
