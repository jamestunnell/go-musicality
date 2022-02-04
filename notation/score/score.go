package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/validation"
)

type Score struct {
	Sections map[string]*section.Section `json:"sections"`
	Program  []string                    `json:"program"`
	Settings map[string]interface{}      `json:"settings"`
}

func New() *Score {
	return &Score{
		Program:  []string{},
		Sections: map[string]*section.Section{},
		Settings: map[string]interface{}{},
	}
}

func (s *Score) Validate() *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	for _, sectionName := range s.Program {
		if _, found := s.Sections[sectionName]; !found {
			err := fmt.Errorf("program references missing section '%s'", sectionName)

			errs = append(errs, err)
		}
	}

	for name, section := range s.Sections {
		if result := section.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %s", result.Context, name)

			results = append(results, result)
		}
	}

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "score",
		Errors:     errs,
		SubResults: results,
	}
}
