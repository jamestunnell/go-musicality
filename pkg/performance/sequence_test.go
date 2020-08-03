package performance_test

import (
	"math/big"
	"testing"

	p "github.com/jamestunnell/go-musicality/pkg/performance"
	"github.com/stretchr/testify/assert"
)

var (
	zero = big.NewRat(0, 1)
	one  = big.NewRat(1, 1)
	// elemA  = &p.Element{Duration: n.Half, Pitch: n.C4, Attack: p.Attack1}
	// elemB1 = &p.Element{Duration: n.One, Pitch: n.D2, Attack: p.Attack2}
	// elemB2 = &p.Element{Duration: n.Quarter, Pitch: n.F3, Attack: p.Attack2}
	// startA = zero
	// startB = one
)

func TestNewSequence(t *testing.T) {
	// c3 := n.NewPitch(3, 0, 0)
	// zeroDurElem := &p.Element{Duration: zero, Pitch: c3, Attack: p.Attack2}
	// nonRestElem := &p.Element{Duration: n.Quarter, Pitch: c3, Attack: p.Attack2}
	// tiedElem := &p.Element{Duration: n.Quarter, Pitch: c3, Attack: p.Attack0}

	// Should fail with no elements
	ns, err := p.NewSequence(zero, []*p.Element{}, p.Separation0)
	assert.Nil(t, ns)
	assert.NotNil(t, err)
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
