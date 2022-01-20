package pitch

import (
	"math"
	"strconv"
)

const (
	// BaseFreq is frequency of C0
	BaseFreq = 16.351597831287414667365624595207
	// SemitonesPerOctave is the number of semitones per octave
	SemitonesPerOctave = 12
)

var semitoneNames = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

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
	p := &Pitch{octave: octave, semitone: semitone}

	p.balance()

	return p
}

func Parse(str string) (*Pitch, error) {
	octave, semitone, err := parse(str)
	if err != nil {
		return nil, err
	}

	return New(octave, semitone), nil
}

func (p *Pitch) Equal(other *Pitch) bool {
	return p.semitone == other.semitone && p.octave == other.octave
}

func (p *Pitch) Diff(other *Pitch) int {
	return p.TotalSemitone() - other.TotalSemitone()
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

func (p *Pitch) Octave() int {
	return p.octave
}

func (p *Pitch) Semitone() int {
	return p.semitone
}

func (p *Pitch) Ratio() float64 {
	totalSem := p.TotalSemitone()

	return math.Pow(2.0, float64(totalSem)/float64(SemitonesPerOctave))
}

func (p *Pitch) Freq() float64 {
	return BaseFreq * p.Ratio()
}

func (p *Pitch) Transpose(semitones int) *Pitch {
	return New(p.octave, p.semitone+semitones)
}

func (p *Pitch) String() string {
	return semitoneNames[p.semitone] + strconv.Itoa(p.octave)
}

func (p *Pitch) TotalSemitone() int {
	return totalSemitone(p.octave, p.semitone)
}

func (p *Pitch) balance() {
	if p.semitone < 0 || p.semitone >= SemitonesPerOctave {
		totalSemitone := p.TotalSemitone()
		p.octave = totalSemitone / SemitonesPerOctave
		p.semitone = totalSemitone - (p.octave * SemitonesPerOctave)
	}
}

func totalSemitone(octave, semitone int) int {
	return octave*SemitonesPerOctave + semitone
}
