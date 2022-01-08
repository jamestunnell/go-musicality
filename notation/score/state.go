package score

import "github.com/jamestunnell/go-musicality/validation"

type State struct {
	Tempo  float64 `json:"tempo"`
	Volume float64 `json:"volume"`
}

func (s *State) Validate() *validation.Result {
	errs := []error{}

	if s.Tempo <= 0.0 {
		errs = append(errs, validation.NewErrNonPositiveFloat("tempo", s.Tempo))
	}

	if s.Volume <= 0.0 {
		errs = append(errs, validation.NewErrNonPositiveFloat("volume", s.Volume))
	}

	if s.Volume > 1.0 {
		errs = append(errs, validation.NewErrNotLessEqualOne("volume", s.Volume))
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "state",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
