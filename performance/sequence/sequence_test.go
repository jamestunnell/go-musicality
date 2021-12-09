package sequence_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/duration"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/sequence"
)

var (
	half = big.NewRat(1, 2)
	one  = big.NewRat(1, 1)
	zero = big.NewRat(0, 1)
)

func TestNewEmpty(t *testing.T) {
	s := sequence.New(zero)

	assert.Equal(t, s.Duration().Cmp(zero), 0)
	assert.Equal(t, s.End().Cmp(zero), 0)
	assert.NoError(t, s.Validate())
	assert.NoError(t, s.Simplify())
}

func TestSequenceInvalid(t *testing.T) {
	e1 := &sequence.Element{Duration: duration.New(1, 4), Pitch: pitch.D4}
	e2 := &sequence.Element{Duration: duration.Zero(), Pitch: pitch.C4}
	e3 := &sequence.Element{Duration: duration.New(1, 2), Pitch: pitch.E4}
	s := sequence.New(half, e1, e2, e3)
	expectedDur := big.NewRat(3, 4)
	expectedEnd := new(big.Rat).Add(half, expectedDur)

	assert.Equal(t, s.Duration().Cmp(expectedDur), 0)
	assert.Equal(t, s.End().Cmp(expectedEnd), 0)
	assert.Error(t, s.Validate())
	assert.Error(t, s.Simplify())
}

func TestValidSequenceValid(t *testing.T) {
	e1 := &sequence.Element{Duration: duration.New(1, 8), Pitch: pitch.D4, Attack: 0.5}
	e2 := &sequence.Element{Duration: duration.New(1, 8), Pitch: pitch.D4, Attack: 0.0}
	e3 := &sequence.Element{Duration: duration.New(1, 2), Pitch: pitch.D4, Attack: 0.5}
	e4 := &sequence.Element{Duration: duration.New(1, 1), Pitch: pitch.E4, Attack: 0.0}
	s := sequence.New(one, e1, e2, e3, e4)
	expectedDur := big.NewRat(7, 4)
	expectedEnd := new(big.Rat).Add(one, expectedDur)

	assert.Equal(t, s.Duration().Cmp(expectedDur), 0)
	assert.Equal(t, s.End().Cmp(expectedEnd), 0)
	assert.NoError(t, s.Validate())

	assert.NoError(t, s.Simplify())

	assert.Equal(t, s.Duration().Cmp(expectedDur), 0)
	assert.Equal(t, s.End().Cmp(expectedEnd), 0)

	require.Len(t, s.Elements, 3)
	assert.True(t, s.Elements[0].Duration.Equal(duration.New(1, 4)))
	assert.True(t, s.Elements[1].Duration.Equal(duration.New(1, 2)))
	assert.True(t, s.Elements[2].Duration.Equal(duration.New(1, 1)))
}

// 	// Should fail without a zero duration element
// 	ns, err = p.NewSequence(zero, []*p.Element{zeroDurElem}, p.Separation0)
// 	assert.Nil(t, ns)
// 	assert.NotNil(t, err)

// 	// Should be okay with a non-rest element
// 	ns, err = p.NewSequence(zero, []*p.Element{nonRestElem}, p.Separation0)
// 	assert.NotNil(t, ns)
// 	assert.Nil(t, err)

// 	// Should simplifiy two tied elements into one element
// 	ns, err = p.NewSequence(zero, []*p.Element{nonRestElem, tiedElem}, p.Separation0)
// 	assert.NotNil(t, ns)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, len(ns.Elements))
// 	assert.True(t, ns.Elements[0].Duration.Equal(nonRestElem.Duration.Add(tiedElem.Duration)))
// }

// func TestSequenceOffsets(t *testing.T) {
// 	testSequenceOffsets(t, makeTestSequenceA(t), []n.NNRational{startA})
// 	testSequenceOffsets(t, makeTestSequenceB(t), []n.NNRational{startB, startB.Add(elemB1.Duration)})
// }

// func testSequenceOffsets(t *testing.T, seq *p.Sequence, expected []n.NNRational) {
// 	actual := seq.Offsets()
// 	if assert.Equal(t, len(expected), len(actual)) {
// 		for i, offset := range expected {
// 			assert.True(t, offset.Equal(actual[i]))
// 		}
// 	}
// }

// func TestSequenceDurationAndEnd(t *testing.T) {
// 	seqA := makeTestSequenceA(t)
// 	assert.True(t, elemA.Duration.Equal(seqA.Duration()))
// 	assert.True(t, startA.Add(elemA.Duration).Equal(seqA.End()))

// 	seqB := makeTestSequenceB(t)
// 	assert.True(t, elemB1.Duration.Add(elemB2.Duration).Equal(seqB.Duration()))
// 	assert.True(t, startB.Add(elemB1.Duration).Add(elemB2.Duration).Equal(seqB.End()))
// }

// func makeTestSequenceA(t *testing.T) *p.Sequence {
// 	seq, err := p.NewSequence(startA, []*p.Element{elemA}, p.Separation0)
// 	assert.NotNil(t, seq)
// 	assert.Nil(t, err)

// 	return seq
// }

// func makeTestSequenceB(t *testing.T) *p.Sequence {
// 	seq, err := p.NewSequence(startB, []*p.Element{elemB1, elemB2}, p.Separation0)
// 	assert.NotNil(t, seq)
// 	assert.Nil(t, err)

// 	return seq
// }
