package centpitch_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/centpitch"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	testBalance(t, 0, 0, 0)
	testBalance(t, 1, 0, 1)
	testBalance(t, 49, 0, 49)
	testBalance(t, 50, 1, -50)
	testBalance(t, 99, 1, -1)
	testBalance(t, 100, 1, 0)
	testBalance(t, 101, 1, 1)
	testBalance(t, 149, 1, 49)
	testBalance(t, 150, 2, -50)
	testBalance(t, -1, 0, -1)
	testBalance(t, -50, 0, -50)
	testBalance(t, -51, -1, 49)
	testBalance(t, -99, -1, 1)
	testBalance(t, -100, -1, 0)
	testBalance(t, -101, -1, -1)
	testBalance(t, -150, -1, -50)
	testBalance(t, -151, -2, 49)
}

func TestCentPitchAllSame(t *testing.T) {
	p1 := centpitch.New(pitch.A2, 0)
	p2 := centpitch.New(pitch.G2, 2*centpitch.CentsPerSemitoneInt)
	p3 := centpitch.New(pitch.B2, -2*centpitch.CentsPerSemitoneInt)

	assert.True(t, p1.Equal(p1))
	assert.True(t, p1.Equal(p2))
	assert.True(t, p1.Equal(p3))

	assert.Equal(t, 0, p1.Diff(p1))
	assert.Equal(t, 0, p1.Diff(p2))
	assert.Equal(t, 0, p1.Diff(p3))

	assert.Equal(t, 0, p1.Compare(p1))
	assert.Equal(t, 0, p1.Compare(p2))
	assert.Equal(t, 0, p1.Compare(p3))

	assert.Equal(t, p1.String(), p2.String())
	assert.Equal(t, p1.String(), p3.String())
}

func TestCentPitchAllDifferent(t *testing.T) {
	p1 := centpitch.New(pitch.A2, 0)
	p2 := centpitch.New(pitch.A2, 3)
	p3 := centpitch.New(pitch.A2, 49)
	p4 := centpitch.New(pitch.A2, 50)
	p5 := centpitch.New(pitch.A2, 99)
	p6 := centpitch.New(pitch.A2, 100)
	p7 := centpitch.New(pitch.A2, -3)
	p8 := centpitch.New(pitch.A2, -50)
	p9 := centpitch.New(pitch.A2, -51)
	p10 := centpitch.New(pitch.A2, -99)
	p11 := centpitch.New(pitch.A2, -100)

	assert.Equal(t, "A2", p1.String())
	assert.Equal(t, "A2+3", p2.String())
	assert.Equal(t, "A2+49", p3.String())
	assert.Equal(t, "Bb2-50", p4.String())
	assert.Equal(t, "Bb2-1", p5.String())
	assert.Equal(t, "Bb2", p6.String())
	assert.Equal(t, "A2-3", p7.String())
	assert.Equal(t, "A2-50", p8.String())
	assert.Equal(t, "Ab2+49", p9.String())
	assert.Equal(t, "Ab2+1", p10.String())
	assert.Equal(t, "Ab2", p11.String())

	less := []*centpitch.CentPitch{p7, p8, p9, p10, p11}
	greater := []*centpitch.CentPitch{p2, p3, p4, p5, p6}

	for _, p := range less {
		assert.False(t, p.Equal(p1))
		assert.Equal(t, -1, p.Compare(p1))
		assert.Less(t, p.Diff(p1), 0)
	}

	for _, p := range greater {
		assert.False(t, p.Equal(p1))
		assert.Equal(t, 1, p.Compare(p1))
		assert.Greater(t, p.Diff(p1), 0)
	}
}

func TestRoundedSemitone(t *testing.T) {
	p1 := centpitch.New(pitch.A2, 0)
	p2 := centpitch.New(pitch.A2, 49)
	p3 := centpitch.New(pitch.A2, 50)
	p4 := centpitch.New(pitch.A2, 99)
	p5 := centpitch.New(pitch.A2, -50)
	p6 := centpitch.New(pitch.A2, -51)
	p7 := centpitch.New(pitch.A2, -99)

	assert.Equal(t, p1.RoundedSemitone(), p2.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()+1, p3.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()+1, p4.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone(), p5.RoundedSemitone())
	assert.Equal(t, p1.RoundedSemitone()-1, p6.RoundedSemitone())
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

func testBalance(t *testing.T, startCentAdjust, expectedSemitoneAdjust, expectedCentAdjust int) {
	t.Run(fmt.Sprintf("balance %d", startCentAdjust), func(t *testing.T) {
		semitoneAdjust, centAdjust := centpitch.Balance(startCentAdjust)

		assert.Equal(t, expectedSemitoneAdjust, semitoneAdjust)
		assert.Equal(t, expectedCentAdjust, centAdjust)
	})
}
