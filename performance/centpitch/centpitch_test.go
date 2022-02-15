package centpitch_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
	"github.com/stretchr/testify/assert"
)

func TestCentPitch(t *testing.T) {
	p1 := centpitch.New(pitch.A2, 0)
	p2 := centpitch.New(pitch.G2, 2*centpitch.CentsPerSemitoneInt)
	p3 := centpitch.New(pitch.A2, 3)
	p4 := centpitch.New(pitch.A2, -3)

	assert.True(t, p1.Equal(p2))
	assert.False(t, p1.Equal(p3))
	assert.False(t, p1.Equal(p4))

	assert.Equal(t, 0, p1.Diff(p2))
	assert.Equal(t, -3, p1.Diff(p3))
	assert.Equal(t, 3, p3.Diff(p1))
	assert.Equal(t, 3, p1.Diff(p4))
	assert.Equal(t, -3, p4.Diff(p1))

	assert.Equal(t, 0, p1.Compare(p2))
	assert.Equal(t, -1, p1.Compare(p3))
	assert.Equal(t, 1, p3.Compare(p1))
	assert.Equal(t, 1, p1.Compare(p4))
	assert.Equal(t, -1, p4.Compare(p1))

	assert.Equal(t, "A2", p1.String())
	assert.Equal(t, "A2", p2.String())
	assert.Equal(t, "A2+3", p3.String())
	assert.Equal(t, "A2-3", p4.String())
}

func TestRoundedSemitone(t *testing.T) {
	p1 := centpitch.New(pitch.A2, 0)
	p2 := centpitch.New(pitch.A2, 49)
	p3 := centpitch.New(pitch.A2, -49)
	p4 := centpitch.New(pitch.A2, 50)
	p5 := centpitch.New(pitch.A2, -50)
	p6 := centpitch.New(pitch.A2, 99)
	p7 := centpitch.New(pitch.A2, -99)

	assert.Equal(t, p1.RoundedSemitone(), p2.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone(), p3.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()+1, p4.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()-1, p5.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()+1, p6.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()-1, p7.RoundedSemitone())
}

func TestFreq(t *testing.T) {
	basePitch := pitch.C3
	startFreq := centpitch.New(basePitch, 0).Freq()

	assert.Equal(t, basePitch.Freq(), startFreq)

	lastFreq := startFreq

	for centAdjust := 1; centAdjust <= 100; centAdjust++ {
		freq := centpitch.New(basePitch, centAdjust).Freq()

		assert.Greater(t, freq, lastFreq)

		lastFreq = freq
	}

	assert.Equal(t, basePitch.Transpose(1).Freq(), lastFreq)

	lastFreq = startFreq

	for centAdjust := -1; centAdjust >= -100; centAdjust-- {
		freq := centpitch.New(basePitch, centAdjust).Freq()

		assert.Less(t, freq, lastFreq)

		lastFreq = freq
	}

	assert.Equal(t, basePitch.Transpose(-1).Freq(), lastFreq)
}
