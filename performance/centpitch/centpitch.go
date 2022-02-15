package centpitch

import (
	"strconv"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

const (
	CentsPerSemitoneInt = 100
	CentsPerSemitoneFlt = 100.0
)

// CentPitch is a cent-adjusted pitch.Pitch.
type CentPitch struct {
	centAdjust int
	*pitch.Pitch
}

func New(p *pitch.Pitch, centAdjust int) *CentPitch {
	p2 := &CentPitch{
		centAdjust: centAdjust,
		Pitch:      p,
	}

	p2.balance()

	return p2
}

func (p *CentPitch) Equal(other *CentPitch) bool {
	return p.Pitch.Equal(other.Pitch) && p.centAdjust == other.centAdjust
}

func (p *CentPitch) Diff(other *CentPitch) int {
	return p.TotalCent() - other.TotalCent()
}

func (p *CentPitch) Compare(other *CentPitch) int {
	diff := p.Diff(other)

	if diff < 0 {
		return -1
	}

	if diff > 0 {
		return 1
	}

	return 0
}

func (p *CentPitch) RoundedSemitone() int {
	totalSem := p.TotalSemitone()

	switch {
	case p.centAdjust <= -50:
		totalSem -= 1
	case p.centAdjust >= 50:
		totalSem += 1
	}

	return totalSem
}

func (p *CentPitch) TotalCent() int {
	return p.TotalSemitone()*CentsPerSemitoneInt + p.centAdjust
}

func (p *CentPitch) Ratio() float64 {
	totalSemitone := float64(p.Pitch.TotalSemitone()) + float64(p.centAdjust)/CentsPerSemitoneFlt

	return pitch.Ratio(totalSemitone)
}

func (p *CentPitch) Freq() float64 {
	return pitch.Freq(p.Ratio())
}

func (p *CentPitch) String() string {
	str := p.Pitch.String()
	if p.centAdjust != 0 {
		str += strconv.Itoa(p.centAdjust)
	}

	return str
}

func (p *CentPitch) balance() {
	if p.centAdjust < -CentsPerSemitoneInt || p.centAdjust >= CentsPerSemitoneInt {
		semitoneAdjust := p.centAdjust / CentsPerSemitoneInt

		p.Pitch = p.Pitch.Transpose(semitoneAdjust)
		p.centAdjust -= semitoneAdjust * CentsPerSemitoneInt
	}
}
