package sequence

import "github.com/jamestunnell/go-musicality/notation/pitch"

const CentsPerSemitone = 100.0

// Pitch is a cent-adjusted pitch.Pitch.
type Pitch struct {
	CentAdjust int
	*pitch.Pitch
}

func NewPitch(p *pitch.Pitch, centAdjust int) *Pitch {
	return &Pitch{
		CentAdjust: centAdjust,
		Pitch:      p,
	}
}

func (p *Pitch) Ratio() float64 {
	totalSemitone := float64(p.Pitch.TotalSemitone()) + float64(p.CentAdjust)/CentsPerSemitone

	return pitch.Ratio(totalSemitone)
}

func (p *Pitch) Freq() float64 {
	return pitch.Freq(p.Ratio())
}
