package measure

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/value"
	"github.com/jamestunnell/go-musicality/validation"
)

type Measure struct {
	Meter          *meter.Meter          `json:"meter"`
	PartNotes      map[string]note.Notes `json:"partNotes"`
	DynamicChanges value.Changes         `json:"dynamicChanges"`
	TempoChanges   value.Changes         `json:"tempoChanges"`
}

const notesDurErrFmt = "total note duration %s does not equal measure duration %s"

func New(met *meter.Meter) *Measure {
	return &Measure{
		Meter:          met,
		PartNotes:      map[string]note.Notes{},
		DynamicChanges: value.Changes{},
		TempoChanges:   value.Changes{},
	}
}

func NewN(n int, met *meter.Meter) []*Measure {
	measures := make([]*Measure, n)

	for i := 0; i < n; i++ {
		measures[i] = New(met)
	}

	return measures
}

func (m *Measure) Duration() *big.Rat {
	return big.NewRat(int64(m.Meter.Numerator), int64(m.Meter.Denominator))
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

		if notesDur.Cmp(dur) != 0 {
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

	if changesResults := validateChanges("dynamic", m.DynamicChanges); len(changesResults) > 0 {
		results = append(results, changesResults...)
	}

	if changesResults := validateChanges("tempo", m.TempoChanges); len(changesResults) > 0 {
		results = append(results, changesResults...)
	}

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "measure",
		Errors:     errs,
		SubResults: results,
	}
}

func validateChanges(changeType string, changes value.Changes) []*validation.Result {
	results := []*validation.Result{}

	for i, change := range changes {
		if result := change.Validate(); result != nil {
			result.Context = fmt.Sprintf("%s %s %d", changeType, result.Context, i)
			results = append(results, result)
		}
	}

	return results
}
