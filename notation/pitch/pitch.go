package pitch

import (
	"fmt"
	"math"
)

const (
	// BaseFreq is frequency of C0
	BaseFreq = 16.351597831287414667365624595207
	// SemitonesPerOctave is the number of semitones per octave
	SemitonesPerOctave = 12
	// Cents per semitone is the number of cents per semitone
	CentsPerSemitone = 100
	// CentsPerOctave is the number of cents per octave
	CentsPerOctave = SemitonesPerOctave * CentsPerSemitone
)

// Pitch represents a musical pitch using octave, semitone, and cent.
type Pitch struct {
	Octave   int `json:"octave"`
	Semitone int `json:"semitone"`
	Cent     int `json:"cent,omitempty"`

	totalCent int
}

// New returns a balance pitch
func New(octave, semitone, cent int) *Pitch {
	p := &Pitch{Octave: octave, Semitone: semitone, Cent: cent}

	if p.IsBalanced() {
		return p
	}

	return p.Balance()
}

func (p *Pitch) Clone() *Pitch {
	return &Pitch{
		Octave:   p.Octave,
		Semitone: p.Semitone,
		Cent:     p.Cent,

		totalCent: p.totalCent,
	}
}

func (p *Pitch) IsBalanced() bool {
	return isIntInRange(p.Semitone, 0, SemitonesPerOctave) &&
		isIntInRange(p.Cent, 0, CentsPerSemitone)
}

func (p *Pitch) Balance() *Pitch {
	remaining := p.TotalCent()

	octave := remaining / CentsPerOctave
	remaining -= octave * CentsPerOctave

	semitone := remaining / CentsPerSemitone
	remaining -= semitone * CentsPerSemitone

	cent := remaining

	return &Pitch{Octave: octave, Semitone: semitone, Cent: cent}
}

func (p *Pitch) TotalCent() int {
	return (p.Octave*SemitonesPerOctave+p.Semitone)*CentsPerSemitone + p.Cent
}

func (p *Pitch) Ratio() float64 {
	return math.Pow(2.0, float64(p.TotalCent())/float64(CentsPerOctave))
}

func (p *Pitch) Freq() float64 {
	return BaseFreq * p.Ratio()
}

func (p *Pitch) Transpose(semitones int) *Pitch {
	return New(p.Octave, p.Semitone+semitones, p.Cent)
}

// Round rounds to the nearest semitone
func (p *Pitch) Round() *Pitch {
	if p.Cent < 50 {
		return New(p.Octave, p.Semitone, 0)
	}

	return New(p.Octave, p.Semitone+1, 0)
}

// TotalSemitoneOffset() converts the (semitone-rounded) pitch to a total semitone offset from C0
func (p *Pitch) TotalSemitoneOffset() int {
	rounded := p.Round()
	return rounded.Octave*SemitonesPerOctave + rounded.Semitone
}

// MIDINote converts the (semitone-rounded) pitch to a MIDI note number.
// If the
func (p *Pitch) MIDINote() (uint8, error) {
	const (
		// minTotalSemitone is the total semitone value of MIDI note 0 (octave below C0)
		minTotalSemitone = -12
		// maxTotalSemitone is the total semitone value of MIDI note 127 (G9)
		maxTotalSemitone = 115
	)

	totalSemitone := p.TotalSemitoneOffset()

	if totalSemitone < minTotalSemitone || totalSemitone > maxTotalSemitone {
		return 0, fmt.Errorf("semitone total %d is outside of MIDI note number range", totalSemitone)
	}

	return uint8(totalSemitone + 12), nil
}

// isIntInRange checks if the given value is in the range [min,max)
func isIntInRange(val, min, max int) bool {
	return (val >= min) && (val < max)
}
