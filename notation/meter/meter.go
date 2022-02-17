package meter

import (
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Meter struct {
	BeatDuration    rat.Rat `json:"beatDuration"`
	BeatsPerMeasure uint64  `json:"beatsPerMeasure"`
}

func New(beatsPerMeasure uint64, beatDuration rat.Rat) *Meter {
	return &Meter{
		BeatDuration:    beatDuration,
		BeatsPerMeasure: beatsPerMeasure,
	}
}

func (m *Meter) MeasureDuration() rat.Rat {
	return m.BeatDuration.MulUint64(uint64(m.BeatsPerMeasure))
}

func (m *Meter) Equal(other *Meter) bool {
	return m.BeatDuration.Equal(other.BeatDuration) && m.BeatsPerMeasure == other.BeatsPerMeasure
}

func (m *Meter) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyNonZeroUInt64("beats per measure", uint64(m.BeatsPerMeasure)); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyPositiveRat("beat duration", m.BeatDuration); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "meter",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
