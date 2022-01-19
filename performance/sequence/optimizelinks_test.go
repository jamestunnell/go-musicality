package sequence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/performance/sequence"
)

func TestOptimizeLinksOneToOne(t *testing.T) {
	ps1 := pitch.NewSet(pitch.C2)
	ps2 := pitch.NewSet(pitch.D5)

	pm := sequence.OptimizeLinks(ps1, ps2)

	assert.Len(t, pm, 1)
	assert.Contains(t, pm, pitch.C2)
	assert.Equal(t, pitch.D5, pm[pitch.C2])
}

func TestOptimizeLinksOneToTwo(t *testing.T) {
	ps1 := pitch.NewSet(pitch.C2)
	ps2 := pitch.NewSet(pitch.D2, pitch.F2)

	pm := sequence.OptimizeLinks(ps1, ps2)

	assert.Len(t, pm, 1)
	assert.Contains(t, pm, pitch.C2)
	assert.Equal(t, pitch.D2, pm[pitch.C2])
}

func TestOptimizeLinksTwoToTwo(t *testing.T) {
	ps1 := pitch.NewSet(pitch.F2, pitch.C2)
	ps2 := pitch.NewSet(pitch.D2, pitch.F2)

	pm := sequence.OptimizeLinks(ps1, ps2)

	assert.Len(t, pm, 2)
	assert.Contains(t, pm, pitch.C2)
	assert.Contains(t, pm, pitch.F2)
	assert.Equal(t, pitch.D2, pm[pitch.C2])
	assert.Equal(t, pitch.F2, pm[pitch.F2])
}

// func TestOptimizeLinksTwoToThree(t *testing.T) {
// 	ps1 := pitch.NewSet(pitch.A2, pitch.G2)
// 	ps2 := pitch.NewSet(pitch.D2, pitch.F2, pitch.C2)

// 	pm := sequence.OptimizeLinks(ps1, ps2)

// 	assert.Len(t, pm, 2)
// 	assert.Contains(t, pm, pitch.G2)
// 	assert.Contains(t, pm, pitch.A2)
// 	assert.Equal(t, pitch.D2, pm[pitch.G2])
// 	assert.Equal(t, pitch.F2, pm[pitch.A2])
// }

func TestScoreLinking(t *testing.T) {
	ps1 := pitch.Pitches{pitch.C2}
	ps2 := pitch.Pitches{pitch.D2}

	assert.Equal(t, 0, sequence.ScoreLinking(ps1, ps1))
	assert.Equal(t, 0, sequence.ScoreLinking(ps2, ps2))
	assert.Equal(t, 2, sequence.ScoreLinking(ps1, ps2))
	assert.Equal(t, 2, sequence.ScoreLinking(ps2, ps1))

	ps1 = pitch.Pitches{pitch.C2, pitch.D2}
	ps2 = pitch.Pitches{pitch.D2, pitch.G2}

	assert.Equal(t, 0, sequence.ScoreLinking(ps1, ps1))
	assert.Equal(t, 0, sequence.ScoreLinking(ps2, ps2))
	assert.Equal(t, 7, sequence.ScoreLinking(ps1, ps2))
	assert.Equal(t, 7, sequence.ScoreLinking(ps2, ps1))

	ps1 = pitch.Pitches{pitch.C3, pitch.E4, pitch.B3}
	ps2 = pitch.Pitches{pitch.G4, pitch.A1, pitch.F1}

	assert.Equal(t, 0, sequence.ScoreLinking(ps1, ps1))
	assert.Equal(t, 0, sequence.ScoreLinking(ps2, ps2))
	assert.Equal(t, 19+31+30, sequence.ScoreLinking(ps1, ps2))
	assert.Equal(t, 19+31+30, sequence.ScoreLinking(ps2, ps1))
}
