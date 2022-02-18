package pitchgen

import "fmt"

// Contains the probabilities for each semitone appearing given a key of C
type CKeyProfile struct {
	semitoneProbs []float64
}

func NewCKeyProfile(semitoneProbs []float64) (*CKeyProfile, error) {
	n := len(semitoneProbs)

	if n != 12 {
		err := fmt.Errorf("wrong number of semitone probabilities: want 12, got %d", n)
		return nil, err
	}

	return &CKeyProfile{semitoneProbs}, nil
}

func (p *CKeyProfile) RotateProbs(n int) []float64 {
	n = (n % 12)

	return append(p.semitoneProbs[:n], p.semitoneProbs[n:]...)
}
