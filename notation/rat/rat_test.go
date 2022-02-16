package rat_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromFloat64(t *testing.T) {
	r := rat.FromFloat64(2.5)

	assert.InDelta(t, 2.5, r.Float64(), 1e-10)
}

func TestFromInt64(t *testing.T) {
	r := rat.New(73, 1)
	r2 := rat.FromInt64(73)

	assert.True(t, r2.Equal(r))
}

func TestFromUint64(t *testing.T) {
	r := rat.New(73, 1)
	r2 := rat.FromUint64(73)

	assert.True(t, r2.Equal(r))
}

func TestMarshalUnmarshal(t *testing.T) {
	testMarshalUnmarshal(t, rat.New(1, 2), `"1/2"`)
	testMarshalUnmarshal(t, rat.New(7, 1), `"7"`)
	testMarshalUnmarshal(t, rat.Zero(), `"0"`)
	testMarshalUnmarshal(t, rat.New(-1, 5), `"-1/5"`)
}

func TestUnmarshalWrongType(t *testing.T) {
	var r rat.Rat

	str := `"not-a-rat"`

	assert.Error(t, json.Unmarshal([]byte(str), &r))
}

func testMarshalUnmarshal(t *testing.T, r rat.Rat, expected string) {
	t.Run(r.String(), func(t *testing.T) {
		d, err := json.Marshal(r)

		require.NoError(t, err)

		assert.Equal(t, expected, string(d))

		var r2 rat.Rat

		require.NoError(t, json.Unmarshal(d, &r2))

		assert.True(t, r2.Equal(r))
	})
}

func TestPositiveZeroNegative(t *testing.T) {
	testPositiveZeroNegative(t, rat.New(1, 2), true, false, false)
	testPositiveZeroNegative(t, rat.Zero(), false, true, false)
	testPositiveZeroNegative(t, rat.New(-1, 2), false, false, true)
}

func TestCompares(t *testing.T) {
	r2 := rat.New(2, 7)
	r1 := r2.Sub(rat.New(1, 100))
	r3 := r2.Add(rat.New(1, 100))

	assert.True(t, r1.Less(r2))
	assert.True(t, r1.LessEqual(r2))
	assert.True(t, r1.LessEqual(r1))
	assert.True(t, r1.Equal(r1))
	assert.True(t, r1.GreaterEqual(r1))
	assert.True(t, r2.GreaterEqual(r1))
	assert.True(t, r2.Greater(r1))

	assert.True(t, r2.Less(r3))
	assert.True(t, r2.LessEqual(r3))
	assert.True(t, r2.LessEqual(r2))
	assert.True(t, r2.Equal(r2))
	assert.True(t, r2.GreaterEqual(r2))
	assert.True(t, r3.GreaterEqual(r2))
	assert.True(t, r3.Greater(r2))
}

func TestMaths(t *testing.T) {
	r1 := rat.New(1, 2)
	r2 := rat.New(3, 4)

	assert.True(t, r1.Add(r2).Equal(rat.New(5, 4)))
	assert.True(t, r1.Sub(r2).Equal(rat.New(-1, 4)))
	assert.True(t, r1.Mul(r2).Equal(rat.New(3, 8)))
	assert.True(t, r1.Div(r2).Equal(rat.New(2, 3)))

	assert.True(t, r1.MulInt64(4).Equal(rat.New(2, 1)))
	assert.True(t, r1.MulUint64(4).Equal(rat.New(2, 1)))
	assert.InDelta(t, 0.25, r1.MulFloat64(0.5).Float64(), 1e-10)
}

func testPositiveZeroNegative(t *testing.T, r rat.Rat, pos, zero, neg bool) {
	assert.Equal(t, pos, r.Positive())
	assert.Equal(t, zero, r.Zero())
	assert.Equal(t, neg, r.Negative())
}
