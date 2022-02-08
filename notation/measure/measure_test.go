package measure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
)

func TestEmpty(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	assert.Nil(t, m.Validate())
	assert.Empty(t, m.PartNotes)
}

func TestInvalidMeter(t *testing.T) {
	m := measure.New(meter.New(0, 2))

	assert.NotNil(t, m.Validate())
}

func TestInvalidPartNote(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(0, 2)),
	}

	assert.NotNil(t, m.Validate())
}

func TestInvalidPartDurs(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(1, 1)),
		note.New(rat.New(1, 1)),
	}

	m.PartNotes["piano"] = []*note.Note{
		note.New(rat.New(1, 4)),
		note.New(rat.New(1, 4)),
	}

	assert.NotNil(t, m.Validate())
}

func TestInvalidDynamicChange(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	// duration is negative
	m.DynamicChanges[rat.Zero()] = change.New(0.5, rat.New(-1, 1))

	assert.NotNil(t, m.Validate())

	// value is out-of-range
	m.DynamicChanges[rat.Zero()] = change.NewImmediate(1.5)
}

func TestInvalidTempoChange(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	// duration is negative
	m.TempoChanges[rat.Zero()] = change.New(100.0, rat.New(-1, 1))

	assert.NotNil(t, m.Validate())

	// value is out-of-range
	m.TempoChanges[rat.Zero()] = change.NewImmediate(-1)
}
