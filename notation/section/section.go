package section

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
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
	DefaultStartDynamic = note.ControlNormal
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

func (s *Section) Duration() rat.Rat {
	dur := rat.Zero()

	for _, m := range s.Measures {
		dur.Accum(m.Duration())
	}

	return dur
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

func (s *Section) PartNames() []string {
	nameMap := map[string]struct{}{}

	for _, m := range s.Measures {
		for name, notes := range m.PartNotes {
			if len(notes) > 0 {
				if _, found := nameMap[name]; !found {
					nameMap[name] = struct{}{}
				}
			}
		}
	}

	names := []string{}

	for name := range nameMap {
		names = append(names, name)
	}

	return names
}

func (s *Section) PartNotes(part string) []*note.Note {
	partNotes := []*note.Note{}

	for _, m := range s.Measures {
		notes, found := m.PartNotes[part]
		if found {
			partNotes = append(partNotes, notes...)
		} else {
			partNotes = append(partNotes, note.New(m.Duration()))
		}
	}

	return partNotes
}

func (s *Section) DynamicChanges(sectionOffset rat.Rat) change.Map {
	return s.gatherChanges(sectionOffset, getDynamicChanges)
}

func (s *Section) TempoChanges(sectionOffset rat.Rat) change.Map {
	return s.gatherChanges(sectionOffset, getTempoChanges)
}

func (s *Section) gatherChanges(sectionOffset rat.Rat, measureChanges func(m *measure.Measure) change.Map) change.Map {
	measureOffset := sectionOffset
	changes := change.Map{}

	for _, m := range s.Measures {
		mChanges := measureChanges(m)

		for offset, c := range mChanges {
			changeOffset := measureOffset.Add(offset)
			change := &change.Change{
				EndValue: c.EndValue,
				Duration: c.Duration,
			}

			changes[changeOffset] = change
		}

		measureOffset.Accum(m.Duration())
	}

	return changes
}

func getDynamicChanges(m *measure.Measure) change.Map {
	return m.DynamicChanges
}

func getTempoChanges(m *measure.Measure) change.Map {
	return m.TempoChanges
}
