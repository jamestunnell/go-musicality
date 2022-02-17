package measure

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/validation"
)

type Measure struct {
	MeterChange    *meter.Meter
	PartNotes      map[string]note.Notes
	DynamicChanges change.Changes
	TempoChanges   change.Changes
}

type measureJSON struct {
	MeterChange    *meter.Meter          `json:"meterChange,omitempty"`
	PartNotes      map[string]note.Notes `json:"partNotes"`
	DynamicChanges changeLiteMap         `json:"dynamicChanges"`
	TempoChanges   changeLiteMap         `json:"tempoChanges"`
}

type changeLiteMap map[string]*changeLite

type changeLite struct {
	EndValue float64 `json:"endValue"`
	Duration rat.Rat `json:"duration"`
}

const notesDurErrFmt = "total note duration %s does not equal measure duration %s"

func New() *Measure {
	return &Measure{
		MeterChange:    nil,
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

func (m *Measure) Validate(measureDur rat.Rat) *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	if m.MeterChange != nil {
		measureDur = m.MeterChange.MeasureDuration()

		if result := m.MeterChange.Validate(); result != nil {
			results = append(results, result)
		}
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

		if !notesDur.Equal(measureDur) {
			err := fmt.Errorf(notesDurErrFmt, notesDur.String(), measureDur.String())

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

func (m *Measure) MarshalJSON() ([]byte, error) {
	dcs, err := toChangeLiteMap(m.DynamicChanges)
	if err != nil {
		err = fmt.Errorf("failed to convert dynamic changes: %w", err)

		return []byte{}, err
	}

	tcs, err := toChangeLiteMap(m.TempoChanges)
	if err != nil {
		err = fmt.Errorf("failed to convert tempo changes: %w", err)

		return []byte{}, err
	}

	mj := &measureJSON{
		MeterChange:    m.MeterChange,
		PartNotes:      m.PartNotes,
		DynamicChanges: dcs,
		TempoChanges:   tcs,
	}

	d, err := json.Marshal(mj)
	if err != nil {
		return []byte{}, err
	}

	return d, nil
}

func (m *Measure) UnmarshalJSON(d []byte) error {
	var mj measureJSON

	err := json.Unmarshal(d, &mj)
	if err != nil {
		return err
	}

	dcs, err := fromChangeLiteMap(mj.DynamicChanges)
	if err != nil {
		err = fmt.Errorf("failed to convert dynamic changes: %w", err)

		return err
	}

	tcs, err := fromChangeLiteMap(mj.TempoChanges)
	if err != nil {
		err = fmt.Errorf("failed to convert tempo changes: %w", err)

		return err
	}

	m.DynamicChanges = dcs
	m.TempoChanges = tcs
	m.MeterChange = mj.MeterChange
	m.PartNotes = mj.PartNotes

	return nil
}

func toChangeLiteMap(changes change.Changes) (changeLiteMap, error) {
	clm := map[string]*changeLite{}

	for _, c := range changes {
		d, err := c.Offset.MarshalJSON()
		if err != nil {
			err = fmt.Errorf("failed to marshal offset: %w", err)

			return changeLiteMap{}, err
		}

		str := strings.Replace(string(d), `"`, "", 2)
		clm[str] = &changeLite{
			EndValue: c.EndValue,
			Duration: c.Duration,
		}
	}

	return clm, nil
}

func fromChangeLiteMap(clm changeLiteMap) (change.Changes, error) {
	changes := change.Changes{}

	for offsetStr, cl := range clm {
		offsetJSONStr := fmt.Sprintf(`"%s"`, offsetStr)

		var offset rat.Rat

		err := json.Unmarshal([]byte(offsetJSONStr), &offset)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal offset: %w", err)

			return change.Changes{}, err
		}

		change := &change.Change{
			Offset:   offset,
			EndValue: cl.EndValue,
			Duration: cl.Duration,
		}
		changes = append(changes, change)
	}

	return changes, nil
}
