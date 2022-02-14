package measure_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestEmpty(t *testing.T) {
	m := measure.New()

	assert.Nil(t, m.Validate(rat.New(1, 1)))
	assert.Empty(t, m.PartNotes)
}

func TestInvalidMeterChange(t *testing.T) {
	m := measure.New()

	m.MeterChange = meter.New(0, rat.New(1, 2))

	assert.NotNil(t, m.Validate(rat.New(1, 1)))
}

func TestInvalidPartNote(t *testing.T) {
	m := measure.New()

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(0, 2)),
	}

	assert.NotNil(t, m.Validate(rat.New(1, 1)))
}

func TestInvalidPartDurs(t *testing.T) {
	m := measure.New()

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(1, 2)),
	}

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(1, 4)),
		note.New(rat.New(1, 4)),
	}

	assert.NotNil(t, m.Validate(rat.New(1, 1)))

	m.MeterChange = meter.TwoFour()

	assert.Nil(t, m.Validate(rat.New(1, 1)))
}

func TestInvalidDynamicChange(t *testing.T) {
	m := measure.New()

	// duration is negative
	m.DynamicChanges = append(m.DynamicChanges, change.New(rat.Zero(), 0.5, rat.New(-1, 1)))

	assert.NotNil(t, m.Validate(rat.New(1, 1)))

	// value is out-of-range
	m.DynamicChanges[0] = change.NewImmediate(rat.Zero(), 1.5)
}

func TestInvalidTempoChange(t *testing.T) {
	m := measure.New()

	// duration is negative
	m.TempoChanges = append(m.TempoChanges, change.New(rat.Zero(), 100.0, rat.New(-1, 1)))

	assert.NotNil(t, m.Validate(rat.New(1, 1)))

	// value is out-of-range
	m.TempoChanges[0] = change.NewImmediate(rat.Zero(), -1)
}

func TestMarshalUnmarshal(t *testing.T) {
	m1 := measure.New()

	m1.DynamicChanges = append(m1.DynamicChanges, change.NewImmediate(rat.Zero(), 1.0))
	m1.DynamicChanges = append(m1.DynamicChanges, change.New(rat.New(1, 2), 0.5, rat.New(2, 1)))
	m1.TempoChanges = append(m1.TempoChanges, change.NewImmediate(rat.Zero(), 100))
	m1.MeterChange = meter.FourFour()

	d, err := json.Marshal(m1)

	require.NoError(t, err)

	var m2 measure.Measure

	require.NoError(t, json.Unmarshal(d, &m2))
}

func TestUnmarshalInvalidJSON(t *testing.T) {
	str := "this-is-not-JSON"

	var m measure.Measure

	require.Error(t, json.Unmarshal([]byte(str), &m))
}
