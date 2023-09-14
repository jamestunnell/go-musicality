package meter

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Meter struct {
	BeatDuration    *big.Rat `json:"beatDuration"`
	BeatsPerMeasure uint64   `json:"beatsPerMeasure"`
}

func New(beatsPerMeasure uint64, beatDuration *big.Rat) *Meter {
	return &Meter{
		BeatDuration:    beatDuration,
		BeatsPerMeasure: beatsPerMeasure,
	}
}

func (m *Meter) String() string {
	num := m.BeatDuration.Num().Uint64() * m.BeatsPerMeasure
	denom := m.BeatDuration.Denom().Uint64()

	return fmt.Sprintf("%d/%d", num, denom)
}

func (m *Meter) MeasureDuration() *big.Rat {
	bpm := new(big.Rat).SetUint64(m.BeatsPerMeasure)

	return rat.Mul(m.BeatDuration, bpm)
}

func (m *Meter) Equal(other *Meter) bool {
	return rat.IsEqual(m.BeatDuration, other.BeatDuration) &&
		m.BeatsPerMeasure == other.BeatsPerMeasure
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
