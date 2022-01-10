package measure

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/validation"
)

type Measure struct {
	Meter     *meter.Meter            `json:"meter"`
	PartNotes map[string][]*note.Note `json:"partNotes"`
	// Changes   map[note.Dur]Change `json:"changes"`
}

func New(met *meter.Meter) *Measure {
	return &Measure{
		Meter:     met,
		PartNotes: map[string][]*note.Note{},
	}
}

func NewN(n int, met *meter.Meter) []*Measure {
	measures := make([]*Measure, n)

	for i := 0; i < n; i++ {
		measures[i] = New(met)
	}

	return measures
}

func (m *Measure) Validate() *validation.Result {
	results := []*validation.Result{}

	if result := m.Meter.Validate(); result != nil {
		results = append(results, result)
	}

	for part, notes := range m.PartNotes {
		partResults := []*validation.Result{}

		for i, note := range notes {
			if result := note.Validate(); result != nil {
				result.Context = fmt.Sprintf("%s %d", result.Context, i)

				partResults = append(partResults, result)
			}
		}

		if len(partResults) > 0 {
			partResult := &validation.Result{
				Context:    fmt.Sprintf("part %s", part),
				Errors:     []error{},
				SubResults: partResults,
			}
			results = append(results, partResult)
		}
	}

	if len(results) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "measure",
		Errors:     []error{},
		SubResults: results,
	}
}
