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
	centAdjust int
	*pitch.Pitch
}

func NewPitch(p *pitch.Pitch, centAdjust int) *Pitch {
	p2 := &Pitch{
		centAdjust: centAdjust,
		Pitch:      p,
	}

	p2.balance()

	return p2
}

func (p *Pitch) Equal(other *Pitch) bool {
	return p.Pitch.Equal(other.Pitch) && p.centAdjust == other.centAdjust
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

func (p *Pitch) RoundedSemitone() int {
	totalSem := p.TotalSemitone()

	switch {
	case p.centAdjust <= -50:
		totalSem -= 1
	case p.centAdjust >= 50:
		totalSem += 1
	}

	return totalSem
}

func (p *Pitch) TotalCent() int {
	return p.TotalSemitone()*CentsPerSemitoneInt + p.centAdjust
}

func (p *Pitch) Ratio() float64 {
	totalSemitone := float64(p.Pitch.TotalSemitone()) + float64(p.centAdjust)/CentsPerSemitoneFlt

	return pitch.Ratio(totalSemitone)
}

func (p *Pitch) Freq() float64 {
	return pitch.Freq(p.Ratio())
}

func (p *Pitch) String() string {
	str := p.Pitch.String()
	if p.centAdjust != 0 {
		str += strconv.Itoa(p.centAdjust)
	}

	return str
}

func (p *Pitch) balance() {
	if p.centAdjust < -CentsPerSemitoneInt || p.centAdjust >= CentsPerSemitoneInt {
		semitoneAdjust := p.centAdjust / CentsPerSemitoneInt

		p.Pitch = p.Pitch.Transpose(semitoneAdjust)
		p.centAdjust -= semitoneAdjust * CentsPerSemitoneInt
	}
}
