package pitch

import "math"

const (
	// BaseFreq is frequency of C0
	BaseFreq = 16.351597831287414667365624595207
	// SemitonesPerOctave is the number of semitones per octave
	SemitonesPerOctave = 12.0
)

func Freq(totalSemitoneOffset float64) float64 {
	return BaseFreq * Ratio(totalSemitoneOffset)
}

func Ratio(totalSemitoneOffset float64) float64 {
	return math.Pow(2.0, totalSemitoneOffset/float64(SemitonesPerOctave))
}
