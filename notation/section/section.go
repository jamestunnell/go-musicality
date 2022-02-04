package section

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/value"
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

func (s *Section) DynamicChanges() value.Changes {
	return s.gatherChanges(getDynamicChanges)
}

func (s *Section) TempoChanges() value.Changes {
	return s.gatherChanges(getTempoChanges)
}

func (s *Section) gatherChanges(measureChanges func(m *measure.Measure) value.Changes) value.Changes {
	measureOffset := big.NewRat(0, 1)
	changes := []*value.Change{}

	for _, m := range s.Measures {
		mChanges := measureChanges(m)
		sort.Sort(mChanges)

		for _, c := range mChanges {
			change := &value.Change{
				Offset:   new(big.Rat).Add(measureOffset, c.Offset),
				EndValue: c.EndValue,
				Duration: c.Duration,
			}
			changes = append(changes, change)
		}

		measureOffset.Add(measureOffset, m.Duration())
	}

	return changes
}

func getDynamicChanges(m *measure.Measure) value.Changes {
	return m.DynamicChanges
}

func getTempoChanges(m *measure.Measure) value.Changes {
	return m.TempoChanges
}
