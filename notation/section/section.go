package section

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/validation"
)

type Section struct {
	StartTempo   float64            `json:"startTempo"`
	StartDynamic float64            `json:"startDynamic"`
	Measures     []*measure.Measure `json:"measures"`
}

type OptFunc func(*Section)

const (
	DefaultStartTempo   = 120.0
	DefaultStartDynamic = 0.0
)

func New(opts ...OptFunc) *Section {
	s := &Section{
		StartTempo:   DefaultStartTempo,
		StartDynamic: DefaultStartDynamic,
		Measures:     []*measure.Measure{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Section) AppendMeasures(n int, met *meter.Meter) {
	s.Measures = append(s.Measures, measure.NewN(n, met)...)
}

func (s *Section) InsertMeasures(n int, met *meter.Meter, idx int) {
	new := s.Measures[:idx]
	new = append(new, measure.NewN(n, met)...)
	new = append(new, s.Measures[idx:]...)

	s.Measures = new
}

func (s *Section) Validate() *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	if err := validation.VerifyPositiveFloat("start tempo", s.StartTempo); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyInRangeFloat("start dynamic", s.StartDynamic, -1.0, 1.0); err != nil {
		errs = append(errs, err)
	}

	for i, m := range s.Measures {
		if result := m.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)

			results = append(results, result)
		}
	}

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "section",
		Errors:     errs,
		SubResults: results,
	}
}
