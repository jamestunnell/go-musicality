package pitchgen

import (
	"math"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/jamestunnell/go-musicality/composition/pitchgen/stats"
	"github.com/jamestunnell/go-musicality/notation/pitch"
)

const (
	// NumOctaves is the number of octaves to support (starting with C0 - B0)
	NumOctaves = 10
	// NumSemitones is the number of octaves to support (starting with C0)
	NumSemitones = NumOctaves * pitch.SemitonesPerOctave
)

var (
	// CMajorBaseKeyProbs contains the probabilities of each octave semitone appearing given a key of C major
	CMajorBaseKeyProbs = NewSemitoneProbs(0.184, 0.001, 0.155, 0.003, 0.191, 0.109, 0.005, 0.214, 0.001, 0.078, 0.004, 0.055)
	// CMajorBaseKeyProbs contains the probabilities of each octave semitone appearing given a key of C minor
	CMinorBaseKeyProbs = NewSemitoneProbs(0.192, 0.005, 0.149, 0.179, 0.002, 0.144, 0.002, 0.201, 0.038, 0.012, 0.053, 0.022)
)

// TemperleyGenerator uses RPK profiles to generate random pitches.
type TemperleyGenerator struct {
	// KeyProbs contains the probabilities of each total semitone offset (from C0) appearing given the current key
	KeyProbs   *SemitoneProbs
	RangeProbs *SemitoneProbs
	last       *pitch.Pitch
	start      *pitch.Pitch
}

type TemperleyOpts struct {
	StartPitch  *pitch.Pitch
	KeySemitone int
	RandSeed    uint64
}

type TemperleyOptSetter func(o *TemperleyOpts)

func NewMajorTemperleyGenerator(optSetters ...TemperleyOptSetter) *TemperleyGenerator {
	return newTemperleyGenerator(CMajorBaseKeyProbs, optSetters...)
}

func NewMinorTemperleyGenerator(optSetters ...TemperleyOptSetter) *TemperleyGenerator {
	return newTemperleyGenerator(CMinorBaseKeyProbs, optSetters...)
}

func TemperleyOptKey(keySemitone int) TemperleyOptSetter {
	return func(o *TemperleyOpts) {
		o.KeySemitone = keySemitone
	}
}

func TemperleyOptStartPitch(p *pitch.Pitch) TemperleyOptSetter {
	return func(o *TemperleyOpts) {
		o.StartPitch = p
	}
}

func TemperleyOptRandSeed(randSeed uint64) TemperleyOptSetter {
	return func(o *TemperleyOpts) {
		o.RandSeed = randSeed
	}
}

func newTemperleyGenerator(baseKeyProbs *SemitoneProbs, optSetters ...TemperleyOptSetter) *TemperleyGenerator {
	opts := &TemperleyOpts{
		RandSeed:    0,
		KeySemitone: 0,
		StartPitch:  nil,
	}

	for _, optSetter := range optSetters {
		optSetter(opts)
	}

	randSrc := rand.NewSource(opts.RandSeed)
	keyProbs := baseKeyProbs.Rotate(opts.KeySemitone)

	centralPitchProfile := distuv.Normal{
		Mu:    float64(56), // semitone offset from C0 - corresponds to Ab4
		Sigma: 3.63,        // stddev - corresponds to variance of about 13.2 semitones
	}

	centralPitchProfile.Src = randSrc

	centralPitchOffset := int(math.Round(centralPitchProfile.Rand()))
	rangeProfile := distuv.Normal{
		Mu:    float64(centralPitchOffset), // semitone offset from C0
		Sigma: 5.39,                        // stddev - corresponds to variance of about 29 semitones
	}

	rangeProfile.Src = randSrc

	rangeProbs := NewSemitoneProbsFromNormal(rangeProfile)

	model := &TemperleyGenerator{
		KeyProbs:   keyProbs,
		RangeProbs: rangeProbs,
		last:       nil,
		start:      opts.StartPitch,
	}

	return model
}

func (pm *TemperleyGenerator) Reset() {
	pm.last = nil
}

func (pm *TemperleyGenerator) NextPitch() *pitch.Pitch {
	if pm.last == nil {
		pm.last = pm.MakeStartingPitch()

		return pm.last
	}

	pm.last = pm.MakeNextPitch(pm.last)

	return pm.last
}

// MakeStartingPitch either uses the given starting pitch in not nil, or uses the
// range and key profiles to determine a random starting pitch.
func (pm *TemperleyGenerator) MakeStartingPitch() *pitch.Pitch {
	if pm.start != nil {
		return pm.start
	}

	probs := CombineAndNormalizeSemitoneProbs(pm.KeyProbs, pm.RangeProbs)

	return pm.makePitch(probs)
}

func (pm *TemperleyGenerator) MakeNextPitch(currentPitch *pitch.Pitch) *pitch.Pitch {
	proximityProfile := distuv.Normal{
		Mu:    float64(currentPitch.TotalSemitone()), // semitone offset from C0
		Sigma: 2.68,                                  // stddev - corresponds to variance of about 7.2 semitones
	}

	proximityProbs := NewSemitoneProbsFromNormal(proximityProfile)
	probs := CombineAndNormalizeSemitoneProbs(pm.KeyProbs, pm.RangeProbs, proximityProbs)

	return pm.makePitch(probs)
}

func (pm *TemperleyGenerator) makePitch(probs *SemitoneProbs) *pitch.Pitch {
	cdf, err := stats.NewCDF(probs.Floats())
	if err != nil {
		log.Fatal().
			Err(err).
			Interface("probs", probs).
			Msg("failed to create CDF")
	}

	i := cdf.Rand()

	return pitch.New(0, i)
}
