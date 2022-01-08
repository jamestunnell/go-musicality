package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/validation"
)

type Score struct {
	Start    *State     `json:"start"`
	Sections []*Section `json:"sections"`
}

func (s *Score) Validate() *validation.Result {
	results := []*validation.Result{}

	if result := s.Start.Validate(); result != nil {
		results = append(results, result)
	}

	for i, section := range s.Sections {
		if result := section.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)

			results = append(results, result)
		}
	}

	if len(results) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "score",
		Errors:     []error{},
		SubResults: results,
	}
}
