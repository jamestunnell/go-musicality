package sequence_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/sequence"
)

func zero() *big.Rat {
	return big.NewRat(0, 1)
}

func TestNewEmpty(t *testing.T) {

	s := sequence.New(zero())

	assert.Equal(t, zero(), s.Duration())
	assert.Equal(t, zero(), s.End())
	assert.Nil(t, s.Validate())
	assert.NoError(t, s.Simplify())
}

func TestSequenceInvalid(t *testing.T) {
	e1 := &sequence.Element{Duration: big.NewRat(1, 4), Pitch: pitch.D4}
	e2 := &sequence.Element{Duration: zero(), Pitch: pitch.C4}
	e3 := &sequence.Element{Duration: big.NewRat(1, 2), Pitch: pitch.E4}
	start := big.NewRat(1, 2)
	s := sequence.New(start, e1, e2, e3)
	expectedDur := big.NewRat(3, 4)
	expectedEnd := new(big.Rat).Add(start, expectedDur)

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())
	assert.NotNil(t, s.Validate())
	assert.Error(t, s.Simplify())
}

func TestValidSequenceValid(t *testing.T) {
	e1 := &sequence.Element{Duration: big.NewRat(1, 8), Pitch: pitch.D4, Attack: note.AttackNormal}
	e2 := &sequence.Element{Duration: big.NewRat(1, 8), Pitch: pitch.D4, Attack: note.AttackMin}
	e3 := &sequence.Element{Duration: big.NewRat(1, 2), Pitch: pitch.D4, Attack: note.AttackNormal}
	e4 := &sequence.Element{Duration: big.NewRat(1, 1), Pitch: pitch.E4, Attack: note.AttackMin}
	start := big.NewRat(1, 1)
	s := sequence.New(start, e1, e2, e3, e4)
	expectedDur := big.NewRat(7, 4)
	expectedEnd := new(big.Rat).Add(start, expectedDur)

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())
	assert.Nil(t, s.Validate())

	assert.NoError(t, s.Simplify())

	assert.Equal(t, expectedDur, s.Duration())
	assert.Equal(t, expectedEnd, s.End())

	require.Len(t, s.Elements, 3)
	assert.Equal(t, big.NewRat(1, 4), s.Elements[0].Duration)
	assert.Equal(t, big.NewRat(1, 2), s.Elements[1].Duration)
	assert.Equal(t, big.NewRat(1, 1), s.Elements[2].Duration)
}

// 	// Should fail without a 0.0 duration element
// 	ns, err = p.NewSequence(0.0, []*p.Element{0.0DurElem}, p.Separation0)
// 	assert.Nil(t, ns)
// 	assert.NotNil(t, err)

// 	// Should be okay with a non-rest element
// 	ns, err = p.NewSequence(0.0, []*p.Element{nonRestElem}, p.Separation0)
// 	assert.NotNil(t, ns)
// 	assert.Nil(t, err)

// 	// Should simplifiy two tied elements into 1.0 element
// 	ns, err = p.NewSequence(0.0, []*p.Element{nonRestElem, tiedElem}, p.Separation0)
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
