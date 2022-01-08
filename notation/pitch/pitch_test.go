package pitch_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

// TestNewPitch tests the NewPitch function, which should return a balanced pitch.
func TestNewPitch(t *testing.T) {
	equivPitches1 := []*pitch.Pitch{
		pitch.New(4, 5, 30),
		pitch.New(5, -7, 30),
		pitch.New(3, 18, -70),
	}

	equivPitches2 := []*pitch.Pitch{
		pitch.New(2, 10, 0),
		pitch.New(3, -2, 0),
		pitch.New(1, 22, 0),
	}

	testNewPitch(t, equivPitches1)
	testNewPitch(t, equivPitches2)
}

// TestPitchIsBalanced tests that a pitch is considered balanced only
// when the semitone and cent are both in ranges [0,11] and [0,99].
// The NewPitch function is intentionally not used, because it always
// creates balanced pitches.
func TestPitchIsBalanced(t *testing.T) {
	balancedPitches := []*pitch.Pitch{
		{Octave: 5, Semitone: 0, Cent: 0},
		{Octave: 3, Semitone: 11, Cent: 65},
		{Octave: 6, Semitone: 4, Cent: 99},
	}

	for _, p := range balancedPitches {
		assert.True(t, p.IsBalanced())
	}

	imbalancedPitches := []*pitch.Pitch{
		{Octave: 5, Semitone: -1, Cent: 0},
		{Octave: 5, Semitone: 7, Cent: -1},
		{Octave: 5, Semitone: 0, Cent: 101},
		{Octave: 3, Semitone: 12, Cent: 0},
		{Octave: 3, Semitone: -4, Cent: -223},
		{Octave: 3, Semitone: 14, Cent: 132},
	}

	for _, p := range imbalancedPitches {
		assert.False(t, p.IsBalanced())
	}
}

// TestPitchIsBalance tests that a pitch which is imbalanced can be turned into
// a balanced pitch, which is equivalent but not equal.
func TestPitchBalance(t *testing.T) {
	imbalancedPitches := []*pitch.Pitch{
		{Octave: 5, Semitone: -1, Cent: 0},
		{Octave: 3, Semitone: 4, Cent: 101},
	}

	for _, p := range imbalancedPitches {
		assert.False(t, p.IsBalanced())

		p2 := p.Balance()

		assert.NotEqual(t, p, p2)
		assert.Equal(t, p.Ratio(), p2.Ratio())
	}
}

func TestPitchRatio(t *testing.T) {
	testCases := map[*pitch.Pitch]float64{
		pitch.New(-1, 0, 0): 0.5,
		pitch.New(0, 0, 0):  1.0,
		pitch.New(1, 0, 0):  2.0,
		pitch.New(2, 0, 0):  4.0,
		pitch.New(0, 1, 0):  math.Pow(2.0, 1.0/12.0),
		pitch.New(0, 3, 0):  math.Pow(2.0, 3.0/12.0),
		pitch.New(0, 0, 1):  math.Pow(2.0, 1.0/1200.0),
		pitch.New(0, 0, 7):  math.Pow(2.0, 7.0/1200.0),
	}

	for p, r := range testCases {
		assert.Equal(t, r, p.Ratio())
	}
}

func TestPitchFreq(t *testing.T) {
	testCases := map[*pitch.Pitch]float64{
		pitch.New(-1, 0, 0): pitch.BaseFreq / 2.0,
		pitch.New(0, 0, 0):  pitch.BaseFreq,
		pitch.New(1, 0, 0):  pitch.BaseFreq * 2.0,
		pitch.New(4, 9, 0):  440.0,
	}

	for p, r := range testCases {
		assert.InDelta(t, r, p.Freq(), 1e-6)
	}
}

func TestPitchTranpose(t *testing.T) {
	startPitch := pitch.New(4, 0, 0)
	testCases := map[int]*pitch.Pitch{
		-1: pitch.New(3, 11, 0),
		0:  pitch.New(4, 0, 0),
		1:  pitch.New(4, 1, 0),
		12: pitch.New(5, 0, 0),
	}

	for semitones, newPitch := range testCases {
		assert.Equal(t, newPitch, startPitch.Transpose(semitones))
	}
}

func TestPitchRound(t *testing.T) {
	testCases := map[*pitch.Pitch]*pitch.Pitch{
		pitch.New(3, 11, 0): pitch.New(3, 11, 0),
		pitch.New(0, 0, 5):  pitch.New(0, 0, 0),
		pitch.New(4, 3, 49): pitch.New(4, 3, 0),
		pitch.New(1, 1, 50): pitch.New(1, 2, 0),
		pitch.New(2, 7, 99): pitch.New(2, 8, 0),
	}

	for startPitch, newPitch := range testCases {
		assert.Equal(t, newPitch, startPitch.Round())
	}
}

func TestPitchMIDINoteInvalid(t *testing.T) {
	testCases := []*pitch.Pitch{
		pitch.New(-2, 0, 0),
		pitch.New(9, 8, 0),
	}

	for _, p := range testCases {
		midiNote, err := p.MIDINote()

		assert.Equal(t, uint(0), midiNote)
		assert.NotNil(t, err)
	}
}

func TestPitchMIDINoteValid(t *testing.T) {
	testCases := map[*pitch.Pitch]uint{
		pitch.New(-1, 0, 0): 0,
		pitch.New(4, 0, 0):  60,
		pitch.New(9, 7, 0):  127,
	}

	for p, expectedMIDINote := range testCases {
		midiNote, err := p.MIDINote()

		assert.Nil(t, err)
		assert.Equal(t, expectedMIDINote, midiNote)
	}
}

// testNewPitch should check that the given pitches, created by NewPitch, are
// all balanced and equal to each other.
func testNewPitch(t *testing.T, equivPitches []*pitch.Pitch) {
	for i, p1 := range equivPitches {
		assert.True(t, p1.IsBalanced())

		for j, p2 := range equivPitches {
			if i != j {
				assert.Equal(t, p1, p2)
			}
		}
	}
}
