package model_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStepPitches(t *testing.T) {
	testStepPitches(t, "same start and end", pitch.C2, pitch.C2, 100)
	// testStepPitches(t, "up two semitones", pitch.C2, pitch.D2,
	// 	pitch.C2, pitch.Db2, pitch.D2)
	// testStepPitches(t, "up five semitones", pitch.C2, pitch.F2,
	// 	pitch.C2, pitch.Db2, pitch.D2, pitch.Eb2, pitch.E2, pitch.F2)
}

func testStepPitches(
	t *testing.T,
	name string,
	start, end *pitch.Pitch,
	centsPerStep int, expected ...*pitch.Pitch) {
	t.Run(name, func(t *testing.T) {
		actual := model.StepPitches(start, end, centsPerStep)
		actual2 := model.StepPitches(end, start, centsPerStep)
		n := len(expected)

		require.Len(t, actual, n)
		require.Len(t, actual2, n)

		for i := 0; i < n; i++ {
			assert.Equal(t, expected[i], actual[i])
			assert.Equal(t, expected[(n-1)-i], actual2[i])
		}
	})
}
