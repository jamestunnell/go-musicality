package pitch_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

// TestNewPitch tests the NewPitch function, which should return a balanced pitch.
func TestNewPitch(t *testing.T) {
	p := pitch.New(7, 5)

	assert.Equal(t, 7, p.Octave())
	assert.Equal(t, 5, p.Semitone())

	equivPitches1 := []*pitch.Pitch{
		pitch.New(4, 5),
		pitch.New(5, -7),
		pitch.New(3, 17),
	}

	equivPitches2 := []*pitch.Pitch{
		pitch.New(2, 10),
		pitch.New(3, -2),
		pitch.New(1, 22),
	}

	equivPitches3 := []*pitch.Pitch{
		pitch.New(0, 0),
		pitch.New(-1, 12),
		pitch.New(1, -12),
	}

	testNewPitch(t, equivPitches1)
	testNewPitch(t, equivPitches2)
	testNewPitch(t, equivPitches3)
}

func TestCompare(t *testing.T) {
	p1 := pitch.C0
	p2 := pitch.D0

	assert.Equal(t, -1, p1.Compare(p2))
	assert.Equal(t, 0, p1.Compare(p1))
	assert.Equal(t, 1, p2.Compare(p1))
}

func TestPitchRatio(t *testing.T) {
	testCases := map[*pitch.Pitch]float64{
		pitch.New(-1, 0): 0.5,
		pitch.New(0, 0):  1.0,
		pitch.New(1, 0):  2.0,
		pitch.New(2, 0):  4.0,
		pitch.New(0, 1):  math.Pow(2.0, 1.0/12.0),
		pitch.New(0, 3):  math.Pow(2.0, 3.0/12.0),
	}

	for p, r := range testCases {
		assert.Equal(t, r, p.Ratio())
	}
}

func TestPitchFreq(t *testing.T) {
	testCases := map[*pitch.Pitch]float64{
		pitch.New(-1, 0): pitch.BaseFreq / 2.0,
		pitch.New(0, 0):  pitch.BaseFreq,
		pitch.New(1, 0):  pitch.BaseFreq * 2.0,
		pitch.New(4, 9):  440.0,
	}

	for p, r := range testCases {
		assert.InDelta(t, r, p.Freq(), 1e-6)
	}
}

func TestPitchTranpose(t *testing.T) {
	startPitch := pitch.New(4, 0)
	testCases := map[int]*pitch.Pitch{
		-1: pitch.New(3, 11),
		0:  pitch.New(4, 0),
		1:  pitch.New(4, 1),
		12: pitch.New(5, 0),
	}

	for semitones, newPitch := range testCases {
		assert.Equal(t, newPitch, startPitch.Transpose(semitones))
	}
}

func TestPitchMarshalUnmarshal(t *testing.T) {
	pitches := pitch.Pitches{
		pitch.C1, pitch.Db1, pitch.D1, pitch.Eb1, pitch.E1, pitch.F1,
		pitch.Gb1, pitch.G1, pitch.Ab1, pitch.A1, pitch.Bb1, pitch.B1,
	}

	for _, p := range pitches {
		d, err := p.MarshalJSON()

		require.NoError(t, err)

		var p2 pitch.Pitch

		err = json.Unmarshal(d, &p2)

		require.NoError(t, err)
		assert.True(t, p2.Equal(p))
	}
}

// func TestPitchMIDINoteInvalid(t *testing.T) {
// 	testCases := []*pitch.Pitch{
// 		pitch.New(-2, 0, 0),
// 		pitch.New(9, 8, 0),
// 	}

// 	for _, p := range testCases {
// 		midiNote, err := p.MIDINote()

// 		assert.Equal(t, uint(0), midiNote)
// 		assert.NotNil(t, err)
// 	}
// }

// func TestPitchMIDINoteValid(t *testing.T) {
// 	testCases := map[*pitch.Pitch]uint{
// 		pitch.New(-1, 0, 0): 0,
// 		pitch.New(4, 0, 0):  60,
// 		pitch.New(9, 7, 0):  127,
// 	}

// 	for p, expectedMIDINote := range testCases {
// 		midiNote, err := p.MIDINote()

// 		assert.Nil(t, err)
// 		assert.Equal(t, expectedMIDINote, midiNote)
// 	}
// }

// testNewPitch should check that the given pitches, created by NewPitch, are
// all balanced and equal to each other.
func testNewPitch(t *testing.T, equivPitches []*pitch.Pitch) {
	for i, p1 := range equivPitches {
		assert.True(t, p1.Semitone() < pitch.SemitonesPerOctave)

		for j, p2 := range equivPitches {
			if i != j {
				assert.Equal(t, p1, p2)
			}
		}
	}
}
