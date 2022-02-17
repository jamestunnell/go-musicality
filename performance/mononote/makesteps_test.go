package mononote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
	"github.com/jamestunnell/go-musicality/performance/mononote"
)

func TestMakeStepsZeroDur(t *testing.T) {
	steps := mononote.MakeSteps(rat.Zero(), pitch.C2, pitch.F2, 25)

	assert.Empty(t, steps)

	steps = mononote.MakeSteps(rat.New(-1, 2), pitch.C2, pitch.F2, 25)

	assert.Empty(t, steps)
}

func TestMakeStepsUp(t *testing.T) {
	steps := mononote.MakeSteps(rat.New(1, 1), pitch.C2, pitch.D2, 25)

	require.Len(t, steps, 8)

	assert.True(t, steps[0].Pitch.Equal(centpitch.New(pitch.C2, 0)))

	for i, pd := range steps {
		assert.True(t, pd.Duration.Equal(rat.New(1, 8)))

		if i > 0 {
			assert.Equal(t, 25, pd.Pitch.Diff(steps[i-1].Pitch))
		}
	}
}

func TestMakeStepsDown(t *testing.T) {
	steps := mononote.MakeSteps(rat.New(1, 1), pitch.D2, pitch.C2, 25)

	require.Len(t, steps, 8)

	assert.True(t, steps[0].Pitch.Equal(centpitch.New(pitch.D2, 0)))

	for i, pd := range steps {
		assert.True(t, pd.Duration.Equal(rat.New(1, 8)))

		if i > 0 {
			assert.Equal(t, -25, pd.Pitch.Diff(steps[i-1].Pitch))
		}
	}
}

func TestMakeStepPitches(t *testing.T) {
	testMakeStepPitches(t, "same start and end", pitch.C2, pitch.C2, 100)
	testMakeStepPitches(t, "up two semitones by 100 cents", pitch.C2, pitch.D2, 100, "C2", "Db2")
	testMakeStepPitches(t, "down two semitones by 100 cents", pitch.D2, pitch.C2, 100, "D2", "Db2")
	testMakeStepPitches(t, "up five semitones by 100 cents", pitch.C2, pitch.F2, 100, "C2", "Db2", "D2", "Eb2", "E2")
	testMakeStepPitches(t, "down five semitones by 100 cents", pitch.F2, pitch.C2, 100, "F2", "E2", "Eb2", "D2", "Db2")
	testMakeStepPitches(t, "up two semitones by 20 cents", pitch.C2, pitch.D2, 20,
		"C2", "C2+20", "C2+40", "Db2-40", "Db2-20", "Db2", "Db2+20", "Db2+40", "D2-40", "D2-20")
	testMakeStepPitches(t, "down two semitones by 20 cents", pitch.D2, pitch.C2, 20,
		"D2", "D2-20", "D2-40", "Db2+40", "Db2+20", "Db2", "Db2-20", "Db2-40", "C2+40", "C2+20")
}

func TestMakePitchDurs(t *testing.T) {
	d := rat.New(1, 4)
	pitches := []*centpitch.CentPitch{
		centpitch.New(pitch.D3, 0),
		centpitch.New(pitch.D3, 10),
		centpitch.New(pitch.D3, 20),
	}

	pds := mononote.MakePitchDurs(d, pitches)

	assert.Len(t, pds, len(pitches))

	for i, pd := range pds {
		assert.True(t, pd.Duration.Equal(d))
		assert.True(t, pd.Pitch.Equal(pitches[i]))
	}
}

func testMakeStepPitches(
	t *testing.T,
	name string,
	start, end *pitch.Pitch,
	centsPerStep int, expected ...string) {
	t.Run(name, func(t *testing.T) {
		actual := mononote.MakeStepPitches(start, end, centsPerStep)
		n := len(expected)

		require.Len(t, actual, n)

		for i := 0; i < n; i++ {
			assert.Equal(t, expected[i], actual[i].String())
		}
	})
}
