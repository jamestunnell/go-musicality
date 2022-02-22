package pitchgen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/composition/pitchgen"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestTemperleyGeneratorDefaults(t *testing.T) {
	g := pitchgen.NewMajorTemperleyGenerator()
	require.NotNil(t, g)

	pitches := pitchgen.MakePitches(16, g)

	verifyPitches(t, pitches, 16)
}

func TestTemperleyGeneratorOpts(t *testing.T) {
	g := pitchgen.NewMinorTemperleyGenerator(
		pitchgen.TemperleyOptStartPitch(pitch.D2),
		pitchgen.TemperleyOptKey(2),
		pitchgen.TemperleyOptRandSeed(123),
	)
	require.NotNil(t, g)

	pitches := pitchgen.MakePitches(22, g)

	verifyPitches(t, pitches, 22)

	assert.True(t, pitches[0].Equal(pitch.D2))

	g.Reset()

	pitches = pitchgen.MakePitches(1, g)

	assert.True(t, pitches[0].Equal(pitch.D2))
}

func verifyPitches(t *testing.T, pitches pitch.Pitches, expectedLen int) {
	assert.Len(t, pitches, expectedLen)

	for _, p := range pitches {
		assert.NotNil(t, p)
	}
}
