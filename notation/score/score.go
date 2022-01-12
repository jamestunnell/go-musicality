package score

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/validation"
)

type Score struct {
	Start    *State                 `json:"start"`
	Sections []*section.Section     `json:"sections"`
	Settings map[string]interface{} `json:"settings"`
}

type OptFunc func(*Score)

const (
	DefaultStartVolume = 1.0
	DefaultStartTempo  = 120.0
)

func OptStartVolume(vol float64) OptFunc {
	return func(s *Score) {
		s.Start.Volume = vol
	}
}

func OptStartTempo(tempo float64) OptFunc {
	return func(s *Score) {
		s.Start.Tempo = tempo
	}
}

func New(opts ...OptFunc) *Score {
	s := &Score{
		Start: &State{
			Tempo:  DefaultStartTempo,
			Volume: DefaultStartVolume,
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
