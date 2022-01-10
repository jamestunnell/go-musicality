package section

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/validation"
)

type Section struct {
	Name     string             `json:"name"`
	Measures []*measure.Measure `json:"measures"`
}

func New(name string) *Section {
	return &Section{
		Name:     name,
		Measures: []*measure.Measure{},
	}
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

	for i, m := range s.Measures {
		if result := m.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)

			results = append(results, result)
		}
	}

	if len(results) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "section",
		Errors:     []error{},
		SubResults: results,
	}
}
