package score

import "github.com/jamestunnell/go-musicality/validation"

type State struct {
	Tempo   float64 `json:"tempo"`
	Dynamic float64 `json:"dynamic"`
}

const (
	DynamicMin = -1.0
	DynamicMax = 1.0
)

func (s *State) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyPositiveFloat("tempo", s.Tempo); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("volume", s.Dynamic, DynamicMin, DynamicMax); err != nil {
		errs = append(errs, err)
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
