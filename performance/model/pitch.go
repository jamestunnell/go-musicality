package model

import (
	"strconv"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

const (
	CentsPerSemitoneInt = 100
	CentsPerSemitoneFlt = 100.0
)

// Pitch is a cent-adjusted pitch.Pitch.
type Pitch struct {
	CentAdjust int
	*pitch.Pitch
}

func NewPitch(p *pitch.Pitch, centAdjust int) *Pitch {
	p2 := &Pitch{
		CentAdjust: centAdjust,
		Pitch:      p,
	}

	p2.balance()

	return p2
}

func (p *Pitch) Equal(other *Pitch) bool {
	return p.Pitch.Equal(other.Pitch) && p.CentAdjust == other.CentAdjust
}

func (p *Pitch) Diff(other *Pitch) int {
	return p.TotalCent() - other.TotalCent()
}

func (p *Pitch) Compare(other *Pitch) int {
	diff := p.Diff(other)

	if diff < 0 {
		return -1
	}

	if diff > 0 {
		return 1
	}

	return 0
}

func (p *Pitch) TotalCent() int {
	return p.TotalSemitone()*CentsPerSemitoneInt + p.CentAdjust
}

func (p *Pitch) Ratio() float64 {
	totalSemitone := float64(p.Pitch.TotalSemitone()) + float64(p.CentAdjust)/CentsPerSemitoneFlt

	return pitch.Ratio(totalSemitone)
}

func (p *Pitch) Freq() float64 {
	return pitch.Freq(p.Ratio())
}

func (p *Pitch) String() string {
	str := p.Pitch.String()
	if p.CentAdjust != 0 {
		str += strconv.Itoa(p.CentAdjust)
	}

	return str
}

func (p *Pitch) balance() {
	if p.CentAdjust < -CentsPerSemitoneInt || p.CentAdjust >= CentsPerSemitoneInt {
		semitoneAdjust := p.CentAdjust / CentsPerSemitoneInt

		p.Pitch = p.Pitch.Transpose(semitoneAdjust)
		p.CentAdjust -= semitoneAdjust * CentsPerSemitoneInt
	}
}
