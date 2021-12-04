package temperley

import (
	"log"
	"math"

	"gonum.org/v1/gonum/stat/distuv"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/pkg/util"
	"github.com/jamestunnell/go-musicality/pkg/util/stats"
)

const (
	// NumOctaves is the number of octaves to support (starting with C0 - B0)
	NumOctaves = 10
	// NumSemitones is the number of octaves to support (starting with C0)
	NumSemitones = NumOctaves * pitch.SemitonesPerOctave
)

var (
	// CMajorBaseKeyProbs contains the probabilities of each octave semitone appearing given a key of C major
	CMajorBaseKeyProbs = []float64{0.184, 0.001, 0.155, 0.003, 0.191, 0.109, 0.005, 0.214, 0.001, 0.078, 0.004, 0.055}
	// CMajorBaseKeyProbs contains the probabilities of each octave semitone appearing given a key of C minor
	CMinorBaseKeyProbs = []float64{0.192, 0.005, 0.149, 0.179, 0.002, 0.144, 0.002, 0.201, 0.038, 0.012, 0.053, 0.022}
	// SemitoneRange covers the range of semitones included in the pitch model
	SemitoneRange = util.NewRange(0, NumSemitones)
)

// PitchModel uses RPK profiles to generate random pitches.
type PitchModel struct {
	// KeyProbs contains the probabilities of each total semitone offset (from C0) appearing given the current key
	KeyProbs     []float64
	RangeProfile distuv.Normal
}

func NewMajorPitchModel(keySemitone uint) (*PitchModel, error) {
	return newPitchModel(keySemitone, CMajorBaseKeyProbs)
}

func NewMinorPitchModel(keySemitone uint) (*PitchModel, error) {
	return newPitchModel(keySemitone, CMinorBaseKeyProbs)
}

func newPitchModel(keySemitone uint, cKeyBaseProbs []float64) (*PitchModel, error) {
	cKeyProfile, err := NewCKeyProfile(cKeyBaseProbs)
	if err != nil {
		return nil, err
	}

	keyBaseProbs := cKeyProfile.RotateProbs(keySemitone)
	keyProbs := make([]float64, NumSemitones)
	for i := 0; i < len(keyProbs); i++ {
		keyProbs[i] = keyBaseProbs[i%12]
	}

	centralPitchProfile := distuv.Normal{
		Mu:    float64(56), // semitone offset from C0 - corresponds to Ab4
		Sigma: 3.63,        // stddev - corresponds to variance of about 13.2 semitones
	}
	centralPitchOffset := int(math.Round(centralPitchProfile.Rand()))
	rangeProfile := distuv.Normal{
		Mu:    float64(centralPitchOffset), // semitone offset from C0
		Sigma: 5.39,                        // stddev - corresponds to variance of about 29 semitones
	}
	model := &PitchModel{
		KeyProbs:     keyProbs,
		RangeProfile: rangeProfile,
	}

	return model, nil
}

// MakeStartingPitch uses the range and key profiles to determine a
// starting pitch
func (pm *PitchModel) MakeStartingPitch() *pitch.Pitch {
	rangeProbs := stats.GetIntProbs(pm.RangeProfile, SemitoneRange)

	return pm.makePitch([][]float64{rangeProbs, pm.KeyProbs})
}

func (pm *PitchModel) MakeNextPitch(currentPitch *pitch.Pitch) *pitch.Pitch {
	proximityProfile := distuv.Normal{
		Mu:    float64(currentPitch.TotalSemitoneOffset()), // semitone offset from C0
		Sigma: 2.68,                                        // stddev - corresponds to variance of about 7.2 semitones
	}

	proximityProbs := stats.GetIntProbs(proximityProfile, SemitoneRange)
	rangeProbs := stats.GetIntProbs(pm.RangeProfile, SemitoneRange)

	return pm.makePitch([][]float64{proximityProbs, rangeProbs, pm.KeyProbs})
}

func (pm *PitchModel) makePitch(probArrays [][]float64) *pitch.Pitch {
	probs, err := stats.CombineAndNormalizeProbs(probArrays)
	if err != nil {
		log.Fatal(err)
	}

	cdf, err := stats.NewCDF(probs)
	if err != nil {
		log.Fatal(err)
	}

	i := cdf.Rand()

	return pitch.New(0, i, 0)
}

func (pm *PitchModel) MakePitches(n uint) []*pitch.Pitch {
	return pm.MakePitchesStartingAt(pm.MakeStartingPitch(), n)
}

func (pm *PitchModel) MakePitchesStartingAt(p *pitch.Pitch, n uint) []*pitch.Pitch {
	switch n {
	case 0:
		return []*pitch.Pitch{}
	case 1:
		return []*pitch.Pitch{p}
	}

	pitches := make([]*pitch.Pitch, n)

	pitches[0] = p

	for i := uint(1); i < n; i++ {
		p = pm.MakeNextPitch(p)
		pitches[i] = p
	}

	return pitches
}
