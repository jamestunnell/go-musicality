package measure

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/validation"
)

type Measure struct {
	Meter          *meter.Meter          `json:"meter"`
	PartNotes      map[string]note.Notes `json:"partNotes"`
	DynamicChanges change.Changes        `json:"dynamicChanges"`
	TempoChanges   change.Changes        `json:"tempoChanges"`
}

const notesDurErrFmt = "total note duration %s does not equal measure duration %s"

func New(met *meter.Meter) *Measure {
	return &Measure{
		Meter:          met,
		PartNotes:      map[string]note.Notes{},
		DynamicChanges: change.Changes{},
		TempoChanges:   change.Changes{},
	}
}

var (
	dynamicRange = &change.MinMaxInclRange{
		Min: note.ControlMin,
		Max: note.ControlMax,
	}
	tempoRange = &change.MinExclRange{
		Min: 0.0,
	}
)

func (m *Measure) Duration() rat.Rat {
	return rat.New(int64(m.Meter.Numerator), int64(m.Meter.Denominator))
}

func (m *Measure) Validate() *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	if result := m.Meter.Validate(); result != nil {
		results = append(results, result)
	}

	dur := m.Duration()
	if err := validation.VerifyPositiveRat("duration", dur); err != nil {
		errs = append(errs, err)
	}

	for part, notes := range m.PartNotes {
		partResults := []*validation.Result{}
		partErrs := []error{}

		for i, note := range notes {
			if result := note.Validate(); result != nil {
				result.Context = fmt.Sprintf("%s %d", result.Context, i)

				partResults = append(partResults, result)
			}
		}

		notesDur := notes.TotalDuration()

		if !notesDur.Equal(dur) {
			err := fmt.Errorf(notesDurErrFmt, notesDur.String(), dur.String())

			partErrs = append(partErrs, err)
		}

		if len(partResults) > 0 || len(partErrs) > 0 {
			partResult := &validation.Result{
				Context:    fmt.Sprintf("part %s", part),
				Errors:     partErrs,
				SubResults: partResults,
			}
			results = append(results, partResult)
		}
	}

	validateChanges := func(changeType string, changes change.Changes, r change.ValueRange) {
		result := changes.Validate(r)
		if result != nil {
			result.Context = fmt.Sprintf("%s %s", changeType, result.Context)

			results = append(results, result)
		}
	}

	validateChanges("dynamic", m.DynamicChanges, dynamicRange)
	validateChanges("tempo", m.TempoChanges, tempoRange)

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "measure",
		Errors:     errs,
		SubResults: results,
	}
}
