package pitch

import (
	"encoding/json"
	"math"
)

const (
	// BaseFreq is frequency of C0
	BaseFreq = 16.351597831287414667365624595207
	// SemitonesPerOctave is the number of semitones per octave
	SemitonesPerOctave = 12
)

// Pitch represents a musical pitch using octave, semitone, and cent.
type Pitch struct {
	octave, semitone int
}

type pitchJSON struct {
	Octave   int `json:"octave"`
	Semitone int `json:"semitone"`
}

// New returns a balance pitch
func New(octave, semitone int) *Pitch {
	if semitone < 0 || semitone >= SemitonesPerOctave {
		octave, semitone = balanced(totalSemitone(octave, semitone))
	}

	return &Pitch{octave: octave, semitone: semitone}
}

func (p *Pitch) MarshalJSON() ([]byte, error) {
	j := pitchJSON{Octave: p.octave, Semitone: p.semitone}

	return json.Marshal(j)
}

func (p *Pitch) UnmarshalJSON(d []byte) error {
	var j pitchJSON

	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}

	if j.Semitone < 0 || j.Semitone > SemitonesPerOctave {
		p.octave, p.semitone = balanced(totalSemitone(j.Octave, j.Semitone))
	} else {
		p.octave = j.Octave
		p.semitone = j.Semitone
	}

	return nil
}

func (p *Pitch) Octave() int {
	return p.octave
}

func (p *Pitch) Semitone() int {
	return p.semitone
}

func totalSemitone(octave, semitone int) int {
	return octave*SemitonesPerOctave + semitone
}

func balanced(totalSemitone int) (octave, semitone int) {
	octave = totalSemitone / SemitonesPerOctave
	semitone = totalSemitone - (octave * SemitonesPerOctave)

	return octave, semitone
}

func (p *Pitch) Ratio() float64 {
	totalSem := totalSemitone(p.octave, p.semitone)

	return math.Pow(2.0, float64(totalSem)/float64(SemitonesPerOctave))
}

func (p *Pitch) Freq() float64 {
	return BaseFreq * p.Ratio()
}

func (p *Pitch) Transpose(semitones int) *Pitch {
	return New(p.octave, p.semitone+semitones)
}
