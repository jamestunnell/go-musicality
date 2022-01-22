package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/validation"
)

type Score struct {
	Start    *State                 `json:"start"`
	Sections []*section.Section     `json:"sections"`
	Settings map[string]interface{} `json:"settings"`
}

type OptFunc func(*Score)

func New(opts ...OptFunc) *Score {
	s := &Score{
		Start: &State{
			Tempo:   DefaultStartTempo,
			Dynamic: DefaultStartDynamic,
		},
		Sections: []*section.Section{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
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
