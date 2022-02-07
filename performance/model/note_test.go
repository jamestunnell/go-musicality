package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/performance/model"
)

var zero = rat.Zero()

func TestNewEmpty(t *testing.T) {

	s := model.NewNote(zero)

	assert.Equal(t, zero, s.Duration())
	assert.Equal(t, zero, s.End())
}

func TestNoteInvalid(t *testing.T) {
	e1 := &model.PitchDur{Duration: rat.New(1, 4), Pitch: model.NewPitch(pitch.D4, 0)}
	e2 := &model.PitchDur{Duration: zero, Pitch: model.NewPitch(pitch.C4, 0)}
	e3 := &model.PitchDur{Duration: rat.New(1, 2), Pitch: model.NewPitch(pitch.E4, 0)}
	start := rat.New(1, 2)
	s := model.NewNote(start, e1, e2, e3)
	expectedDur := rat.New(3, 4)
	expectedEnd := start.Add(expectedDur)

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())
}

func TestValidNoteValid(t *testing.T) {
	e1 := &model.PitchDur{Duration: rat.New(1, 8), Pitch: model.NewPitch(pitch.D4, 0)}
	e2 := &model.PitchDur{Duration: rat.New(1, 8), Pitch: model.NewPitch(pitch.D4, 0)}
	e3 := &model.PitchDur{Duration: rat.New(1, 2), Pitch: model.NewPitch(pitch.D4, 0)}
	e4 := &model.PitchDur{Duration: rat.New(1, 1), Pitch: model.NewPitch(pitch.E4, 0)}
	start := rat.New(1, 1)
	s := model.NewNote(start, e1, e2, e3, e4)
	expectedDur := rat.New(7, 4)
	expectedEnd := start.Add(expectedDur)

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())

	s.Simplify()

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())

	require.Len(t, s.PitchDurs, 2)
	assert.Equal(t, rat.New(3, 4), s.PitchDurs[0].Duration)
	assert.Equal(t, rat.New(1, 1), s.PitchDurs[1].Duration)
}

// 	// Should fail without a 0.0 duration PitchDur
// 	ns, err = p.NewNote(0.0, []*p.PitchDur{0.0DurElem}, p.Separation0)
// 	assert.Nil(t, ns)
// 	assert.NotNil(t, err)

// 	// Should be okay with a non-rest PitchDur
// 	ns, err = p.NewNote(0.0, []*p.PitchDur{nonRestElem}, p.Separation0)
// 	assert.NotNil(t, ns)
// 	assert.Nil(t, err)

// 	// Should simplifiy two tied PitchDurs into 1.0 PitchDur
// 	ns, err = p.NewNote(0.0, []*p.PitchDur{nonRestElem, tiedElem}, p.Separation0)
// 	assert.NotNil(t, ns)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, len(ns.PitchDurs))
// 	assert.True(t, ns.PitchDurs[0].Duration.Equal(nonRestElem.Duration.Add(tiedElem.Duration)))
// }

// func TestNoteOffsets(t *testing.T) {
// 	testNoteOffsets(t, makeTestNoteA(t), []n.NNRational{startA})
// 	testNoteOffsets(t, makeTestNoteB(t), []n.NNRational{startB, startB.Add(elemB1.Duration)})
// }

// func testNoteOffsets(t *testing.T, seq *p.Note, expected []n.NNRational) {
// 	actual := seq.Offsets()
// 	if assert.Equal(t, len(expected), len(actual)) {
// 		for i, offset := range expected {
// 			assert.True(t, offset.Equal(actual[i]))
// 		}
// 	}
// }

// func TestNoteDurationAndEnd(t *testing.T) {
// 	seqA := makeTestNoteA(t)
// 	assert.True(t, elemA.Duration.Equal(seqA.Duration()))
// 	assert.True(t, startA.Add(elemA.Duration).Equal(seqA.End()))

// 	seqB := makeTestNoteB(t)
// 	assert.True(t, elemB1.Duration.Add(elemB2.Duration).Equal(seqB.Duration()))
// 	assert.True(t, startB.Add(elemB1.Duration).Add(elemB2.Duration).Equal(seqB.End()))
// }

// func makeTestNoteA(t *testing.T) *p.Note {
// 	seq, err := p.NewNote(startA, []*p.PitchDur{elemA}, p.Separation0)
// 	assert.NotNil(t, seq)
// 	assert.Nil(t, err)

// 	return seq
// }

// func makeTestNoteB(t *testing.T) *p.Note {
// 	seq, err := p.NewNote(startB, []*p.PitchDur{elemB1, elemB2}, p.Separation0)
// 	assert.NotNil(t, seq)
// 	assert.Nil(t, err)

// 	return seq
// }
