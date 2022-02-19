package pitchgen

import (
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"gonum.org/v1/gonum/stat/distuv"
)

type SemitoneProbs struct {
	C, Db, D, Eb, E, F, Gb, G, Ab, A, Bb, B float64

	ptrs []*float64
}

func NewSemitoneProbs(c, db, d, eb, e, f, gb, g, ab, a, bb, b float64) *SemitoneProbs {
	return newSemitoneProbs(c, db, d, eb, e, f, gb, g, ab, a, bb, b)
}

func NewSemitoneProbsAllOnes() *SemitoneProbs {
	return newSemitoneProbs(1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0)
}

func NewSemitoneProbsFromNormal(dist distuv.Normal) *SemitoneProbs {
	return newSemitoneProbs(
		dist.Prob(float64(0)),
		dist.Prob(float64(1)),
		dist.Prob(float64(2)),
		dist.Prob(float64(3)),
		dist.Prob(float64(4)),
		dist.Prob(float64(5)),
		dist.Prob(float64(6)),
		dist.Prob(float64(7)),
		dist.Prob(float64(8)),
		dist.Prob(float64(9)),
		dist.Prob(float64(10)),
		dist.Prob(float64(11)),
	)
}

func newSemitoneProbs(c, db, d, eb, e, f, gb, g, ab, a, bb, b float64) *SemitoneProbs {
	sp := &SemitoneProbs{
		C:  c,
		Db: db,
		D:  d,
		Eb: eb,
		E:  e,
		F:  f,
		Gb: gb,
		G:  g,
		Ab: ab,
		A:  a,
		Bb: bb,
		B:  b,
	}

	sp.ptrs = []*float64{&sp.C, &sp.Db, &sp.D, &sp.Eb, &sp.E, &sp.F, &sp.Gb, &sp.G, &sp.Ab, &sp.A, &sp.Bb, &sp.B}

	return sp
}

func (sp *SemitoneProbs) Rotate(n int) *SemitoneProbs {
	n = (n % pitch.SemitonesPerOctave)

	spNew := NewSemitoneProbsAllOnes()

	for i := 0; i < pitch.SemitonesPerOctave; i++ {
		srcIdx := (i + n) % pitch.SemitonesPerOctave

		*spNew.ptrs[i] = *sp.ptrs[srcIdx]
	}

	return spNew
}

func (sp *SemitoneProbs) Mul(other *SemitoneProbs) {
	for i := 0; i < pitch.SemitonesPerOctave; i++ {
		*sp.ptrs[i] *= *other.ptrs[i]
	}
}

func (sp *SemitoneProbs) Normalize() {
	total := 0.0

	for i := 0; i < pitch.SemitonesPerOctave; i++ {
		total += *sp.ptrs[i]
	}

	mul := 1.0 / total

	for i := 0; i < pitch.SemitonesPerOctave; i++ {
		*sp.ptrs[i] *= mul
	}
}

func (sp *SemitoneProbs) Floats() []float64 {
	return []float64{sp.C, sp.Db, sp.D, sp.Eb, sp.E, sp.F, sp.Gb, sp.G, sp.Ab, sp.A, sp.Bb, sp.B}
}

func CombineAndNormalizeSemitoneProbs(first *SemitoneProbs, more ...*SemitoneProbs) *SemitoneProbs {
	probs := NewSemitoneProbsAllOnes()

	probs.Mul(first)

	for _, another := range more {
		probs.Mul(another)
	}

	probs.Normalize()

	return probs
}
