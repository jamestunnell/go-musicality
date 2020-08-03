package notation_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/pkg/notation"
)

// TestNewPitch tests the NewPitch function, which should return a balanced pitch.
func TestNewPitch(t *testing.T) {
	equivPitches1 := []*notation.Pitch{
		notation.NewPitch(4, 5, 30),
		notation.NewPitch(5, -7, 30),
		notation.NewPitch(3, 18, -70),
	}

	equivPitches2 := []*notation.Pitch{
		notation.NewPitch(2, 10, 0),
		notation.NewPitch(3, -2, 0),
		notation.NewPitch(1, 22, 0),
	}

	testNewPitch(t, equivPitches1)
	testNewPitch(t, equivPitches2)
}

// TestPitchIsBalanced tests that a pitch is considered balanced only
// when the semitone and cent are both in ranges [0,11] and [0,99].
// The NewPitch function is intentionally not used, because it always
// creates balanced pitches.
func TestPitchIsBalanced(t *testing.T) {
	balancedPitches := []*notation.Pitch{
		&notation.Pitch{Octave: 5, Semitone: 0, Cent: 0},
		&notation.Pitch{Octave: 3, Semitone: 11, Cent: 65},
		&notation.Pitch{Octave: 6, Semitone: 4, Cent: 99},
	}

	for _, p := range balancedPitches {
		assert.True(t, p.IsBalanced())
	}

	imbalancedPitches := []*notation.Pitch{
		&notation.Pitch{Octave: 5, Semitone: -1, Cent: 0},
		&notation.Pitch{Octave: 5, Semitone: 7, Cent: -1},
		&notation.Pitch{Octave: 5, Semitone: 0, Cent: 101},
		&notation.Pitch{Octave: 3, Semitone: 12, Cent: 0},
		&notation.Pitch{Octave: 3, Semitone: -4, Cent: -223},
		&notation.Pitch{Octave: 3, Semitone: 14, Cent: 132},
	}

	for _, p := range imbalancedPitches {
		assert.False(t, p.IsBalanced())
	}
}

// TestPitchIsBalance tests that a pitch which is imbalanced can be turned into
// a balanced pitch, which is equivalent but not equal.
func TestPitchBalance(t *testing.T) {
	imbalancedPitches := []*notation.Pitch{
		&notation.Pitch{Octave: 5, Semitone: -1, Cent: 0},
		&notation.Pitch{Octave: 3, Semitone: 4, Cent: 101},
	}

	for _, p := range imbalancedPitches {
		assert.False(t, p.IsBalanced())

		p2 := p.Balance()

		assert.NotEqual(t, p, p2)
		assert.Equal(t, p.Ratio(), p2.Ratio())
	}
}

func TestPitchRatio(t *testing.T) {
	testCases := map[*notation.Pitch]float64{
		notation.NewPitch(-1, 0, 0): 0.5,
		notation.NewPitch(0, 0, 0):  1.0,
		notation.NewPitch(1, 0, 0):  2.0,
		notation.NewPitch(2, 0, 0):  4.0,
		notation.NewPitch(0, 1, 0):  math.Pow(2.0, 1.0/12.0),
		notation.NewPitch(0, 3, 0):  math.Pow(2.0, 3.0/12.0),
		notation.NewPitch(0, 0, 1):  math.Pow(2.0, 1.0/1200.0),
		notation.NewPitch(0, 0, 7):  math.Pow(2.0, 7.0/1200.0),
	}

	for p, r := range testCases {
		assert.Equal(t, r, p.Ratio())
	}
}

func TestPitchFreq(t *testing.T) {
	testCases := map[*notation.Pitch]float64{
		notation.NewPitch(-1, 0, 0): notation.BaseFreq / 2.0,
		notation.NewPitch(0, 0, 0):  notation.BaseFreq,
		notation.NewPitch(1, 0, 0):  notation.BaseFreq * 2.0,
		notation.NewPitch(4, 9, 0):  440.0,
	}

	for p, r := range testCases {
		assert.Equal(t, r, p.Freq())
	}
}

func TestPitchTranpose(t *testing.T) {
	startPitch := notation.NewPitch(4, 0, 0)
	testCases := map[int]*notation.Pitch{
		-1: notation.NewPitch(3, 11, 0),
		0:  notation.NewPitch(4, 0, 0),
		1:  notation.NewPitch(4, 1, 0),
		12: notation.NewPitch(5, 0, 0),
	}

	for semitones, newPitch := range testCases {
		assert.Equal(t, newPitch, startPitch.Transpose(semitones))
	}
}

func TestPitchRound(t *testing.T) {
	testCases := map[*notation.Pitch]*notation.Pitch{
		notation.NewPitch(3, 11, 0): notation.NewPitch(3, 11, 0),
		notation.NewPitch(0, 0, 5):  notation.NewPitch(0, 0, 0),
		notation.NewPitch(4, 3, 49): notation.NewPitch(4, 3, 0),
		notation.NewPitch(1, 1, 50): notation.NewPitch(1, 2, 0),
		notation.NewPitch(2, 7, 99): notation.NewPitch(2, 8, 0),
	}

	for startPitch, newPitch := range testCases {
		assert.Equal(t, newPitch, startPitch.Round())
	}
}

func TestPitchMIDINoteInvalid(t *testing.T) {
	testCases := []*notation.Pitch{
		notation.NewPitch(-2, 0, 0),
		notation.NewPitch(9, 8, 0),
	}

	for _, p := range testCases {
		midiNote, err := p.MIDINote()

		assert.Equal(t, uint(0), midiNote)
		assert.NotNil(t, err)
	}
}

func TestPitchMIDINoteValid(t *testing.T) {
	testCases := map[*notation.Pitch]uint{
		notation.NewPitch(-1, 0, 0): 0,
		notation.NewPitch(4, 0, 0):  60,
		notation.NewPitch(9, 7, 0):  127,
	}

	for p, expectedMIDINote := range testCases {
		midiNote, err := p.MIDINote()

		assert.Nil(t, err)
		assert.Equal(t, expectedMIDINote, midiNote)
	}
}

// testNewPitch should check that the given pitches, created by NewPitch, are
// all balanced and equal to each other.
func testNewPitch(t *testing.T, equivPitches []*notation.Pitch) {
	for i, p1 := range equivPitches {
		assert.True(t, p1.IsBalanced())

		for j, p2 := range equivPitches {
			if i != j {
				assert.Equal(t, p1, p2)
			}
		}
	}
}
