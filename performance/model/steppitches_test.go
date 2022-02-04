package sequence_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/sequence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntRangeAsc(t *testing.T) {
	assert.Equal(t, []int{}, sequence.IntRangeAsc(5, 1))
	assert.Equal(t, []int{5}, sequence.IntRangeAsc(5, 5))
	assert.Equal(t, []int{0, 1, 2, 3}, sequence.IntRangeAsc(0, 3))
	assert.Equal(t, []int{-3, -2, -1, 0}, sequence.IntRangeAsc(-3, 0))
}

func TestIntRangeDesc(t *testing.T) {
	assert.Equal(t, []int{}, sequence.IntRangeDesc(1, 5))
	assert.Equal(t, []int{5}, sequence.IntRangeDesc(5, 5))
	assert.Equal(t, []int{3, 2, 1, 0}, sequence.IntRangeDesc(3, 0))
	assert.Equal(t, []int{0, -1, -2, -3}, sequence.IntRangeDesc(0, -3))
}

func TestStepPitches(t *testing.T) {
	testStepPitches(t, "same start and end", pitch.C2, pitch.C2,
		pitch.C2)
	testStepPitches(t, "up two semitones", pitch.C2, pitch.D2,
		pitch.C2, pitch.Db2, pitch.D2)
	testStepPitches(t, "up five semitones", pitch.C2, pitch.F2,
		pitch.C2, pitch.Db2, pitch.D2, pitch.Eb2, pitch.E2, pitch.F2)
}

func testStepPitches(t *testing.T, name string, start, end *pitch.Pitch, expected ...*pitch.Pitch) {
	t.Run(name, func(t *testing.T) {
		actual := sequence.StepPitches(start, end)
		actual2 := sequence.StepPitches(end, start)
		n := len(expected)

		require.Len(t, actual, n)
		require.Len(t, actual2, n)

		for i := 0; i < n; i++ {
			assert.Equal(t, expected[i], actual[i])
			assert.Equal(t, expected[(n-1)-i], actual2[i])
		}
	})
}
