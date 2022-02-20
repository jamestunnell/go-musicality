package pitchgen

import "gonum.org/v1/gonum/stat/distuv"

// Probabilities of each pitch, starting at 0 -> C0. Not normalized.
type PitchProbs [NumSemitones]float64

func NewPitchProbsFromNormal(dist distuv.Normal) PitchProbs {
	probs := PitchProbs{}
	for i := 0; i < NumSemitones; i++ {
		probs[i] = dist.Prob(float64(i))
	}

	return probs
}

func CombineAndNormalizePitchProbs(first PitchProbs, more ...PitchProbs) PitchProbs {
	probs := PitchProbs{}

	for i := 0; i < NumSemitones; i++ {
		probs[i] = first[i]
	}

	for _, another := range more {
		for i := 0; i < NumSemitones; i++ {
			probs[i] *= another[i]
		}
	}

	sum := 0.0
	for _, prob := range probs {
		sum += prob
	}

	mul := 1.0 / sum

	for i := 0; i < NumSemitones; i++ {
		probs[i] *= mul
	}

	return probs
}
