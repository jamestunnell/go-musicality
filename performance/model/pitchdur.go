package model

import "github.com/jamestunnell/go-musicality/notation/rat"

type PitchDur struct {
	Duration rat.Rat
	Pitch    *Pitch
}

func NewPitchDur(p *Pitch, dur rat.Rat) *PitchDur {
	return &PitchDur{
		Pitch:    p,
		Duration: dur,
	}
}

// func (e *PitchDur) Validate() *validation.Result {
// 	errs := []error{}

// 	if err := validation.VerifyPositiveRat("duration", e.Duration); err != nil {
// 		errs = append(errs, err)
// 	}

// 	if len(errs) == 0 {
// 		return nil
// 	}

// 	return &validation.Result{
// 		Context:    "PitchDur",
// 		Errors:     errs,
// 		SubResults: []*validation.Result{},
// 	}
// }
