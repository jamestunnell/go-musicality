package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/note"
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

func (s *Score) PartNames() []string {
	nameMap := map[string]struct{}{}

	for _, section := range s.Sections {
		for _, m := range section.Measures {
			for name, notes := range m.PartNotes {
				if len(notes) > 0 {
					if _, found := nameMap[name]; !found {
						nameMap[name] = struct{}{}
					}
				}
			}
		}
	}

	names := make([]string, len(nameMap))
	i := 0

	for name := range nameMap {
		names[i] = name

		i++
	}

	return names
}

func (s *Score) PartNotes(part string) []*note.Note {
	partNotes := []*note.Note{}

	for _, section := range s.Sections {
		for _, m := range section.Measures {
			notes, found := m.PartNotes[part]
			if found {
				partNotes = append(partNotes, notes...)
			} else {
				partNotes = append(partNotes, note.New(m.Duration()))
			}
		}
	}

	return partNotes
}
