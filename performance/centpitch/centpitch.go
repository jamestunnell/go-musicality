package centpitch

import (
	"strconv"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

const (
	MinCentAdjust       = -50
	MaxCentAdjust       = 49
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
	return p.Pitch.TotalSemitone()
}

func (p *CentPitch) TotalCent() int {
	return p.Pitch.TotalSemitone()*CentsPerSemitoneInt + p.centAdjust
}

func (p *CentPitch) TotalSemitone() float64 {
	return float64(p.TotalCent()) / CentsPerSemitoneFlt
}

// func (p *CentPitch) Ratio() float64 {
// 	return pitch.Ratio(p.TotalSemitone())
// }

func (p *CentPitch) Freq() float64 {
	return pitch.Freq(p.TotalSemitone())
}

func (p *CentPitch) String() string {
	str := p.Pitch.String()
	if p.centAdjust != 0 {
		if p.centAdjust > 0 {
			str += "+"
		}
		str += strconv.Itoa(p.centAdjust)
	}

	return str
}

func (p *CentPitch) balance() {
	semitoneAdjust, centAdjust := Balance(p.centAdjust)

	if semitoneAdjust != 0 {
		p.Pitch = p.Pitch.Transpose(semitoneAdjust)
	}

	p.centAdjust = centAdjust
}

func Balance(startCentAdjust int) (semitoneAdjust, centAdjust int) {
	semitoneAdjust = 0
	centAdjust = startCentAdjust

	// An initial adjustment should bring the cent adjust to within [-99,99]
	if centAdjust <= -CentsPerSemitoneInt || centAdjust >= CentsPerSemitoneInt {
		semitoneAdjust = centAdjust / CentsPerSemitoneInt
		centAdjust -= semitoneAdjust * CentsPerSemitoneInt
	}

	// Further adjustment to bring it into [-50,49]
	if centAdjust < MinCentAdjust {
		semitoneAdjust -= 1
		centAdjust += CentsPerSemitoneInt
	} else if centAdjust > MaxCentAdjust {
		semitoneAdjust += 1
		centAdjust -= CentsPerSemitoneInt
	}

	return
}
