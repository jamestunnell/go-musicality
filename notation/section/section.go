package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/validation"
)

type Section struct {
	Name     string             `json:"name"`
	Measures []*measure.Measure `json:"measures"`
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
