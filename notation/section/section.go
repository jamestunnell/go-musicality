package section

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/validation"
)

type Section struct {
	StartTempo float64            `json:"startTempo"`
	StartMeter *meter.Meter       `json:"startMeter"`
	Measures   []*measure.Measure `json:"measures"`
}

type OptFunc func(*Section)

const (
	DefaultStartTempo = 120.0
)

func New(opts ...OptFunc) *Section {
	s := &Section{
		StartTempo: DefaultStartTempo,
		StartMeter: meter.FourFour(),
		Measures:   []*measure.Measure{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Section) Duration() *big.Rat {
	dur := rat.Zero()
	mDur := s.StartMeter.MeasureDuration()

	for _, m := range s.Measures {
		if m.MeterChange != nil {
			mDur = m.MeterChange.MeasureDuration()
		}

		dur = rat.Add(dur, mDur)
	}

	return dur
}

func (s *Section) Validate() *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	if err := validation.VerifyPositiveFloat("start tempo", s.StartTempo); err != nil {
		errs = append(errs, err)
	}

	if result := s.StartMeter.Validate(); result != nil {
		result.Context = "start " + result.Context

		results = append(results, result)
	} else {
		mDur := s.StartMeter.MeasureDuration()

		for i, m := range s.Measures {
			if result := m.Validate(mDur); result != nil {
				result.Context = fmt.Sprintf("%s %d", result.Context, i)

				results = append(results, result)
			}

			if m.MeterChange != nil {
				mDur = m.MeterChange.MeasureDuration()
			}
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
	mDur := s.StartMeter.MeasureDuration()

	for _, m := range s.Measures {
		if m.MeterChange != nil {
			mDur = m.MeterChange.MeasureDuration()
		}

		notes, found := m.PartNotes[part]
		if found {
			partNotes = append(partNotes, notes...)
		} else {
			partNotes = append(partNotes, note.New(mDur))
		}
	}

	return partNotes
}

func (s *Section) BeatDurChanges(sectionOffset *big.Rat) change.Changes {
	return s.gatherChanges(sectionOffset, getBeatDurChanges)
}

func (s *Section) TempoChanges(sectionOffset *big.Rat) change.Changes {
	return s.gatherChanges(sectionOffset, getTempoChanges)
}

func (s *Section) gatherChanges(sectionOffset *big.Rat, measureChanges func(m *measure.Measure) change.Changes) change.Changes {
	measureOffset := sectionOffset
	changes := change.Changes{}
	mDur := s.StartMeter.MeasureDuration()

	for _, m := range s.Measures {
		if m.MeterChange != nil {
			mDur = m.MeterChange.MeasureDuration()
		}

		mChanges := measureChanges(m)

		for _, c := range mChanges {
			change := &change.Change{
				Offset:   rat.Add(measureOffset, c.Offset),
				EndValue: c.EndValue,
				Duration: c.Duration,
			}

			changes = append(changes, change)
		}

		measureOffset = rat.Add(measureOffset, mDur)
	}

	return changes
}

func getBeatDurChanges(m *measure.Measure) change.Changes {
	if m.MeterChange == nil {
		return change.Changes{}
	}

	beatDur, _ := m.MeterChange.BeatDuration.Float64()

	c := change.NewImmediate(rat.Zero(), beatDur)

	return change.Changes{c}
}

func getTempoChanges(m *measure.Measure) change.Changes {
	return m.TempoChanges
}
