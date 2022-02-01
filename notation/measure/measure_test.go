package measure_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
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
		note.New(big.NewRat(0, 2)),
	}

	assert.NotNil(t, m.Validate())
}

func TestInvalidPartDurs(t *testing.T) {
	m := measure.New(meter.New(4, 4))

	m.PartNotes["piano"] = []*note.Note{
		note.New(big.NewRat(1, 1)),
		note.New(big.NewRat(1, 1)),
	}

	m.PartNotes["piano"] = []*note.Note{
		note.New(big.NewRat(1, 4)),
		note.New(big.NewRat(1, 4)),
	}

	assert.NotNil(t, m.Validate())
}

// func TestValidateJSON(t *testing.T) {
// 	testValidateJSONFails(t, "{}")
// 	testValidateJSONFails(t, `{"meter":"4/4"}`)

// 	testValidateJSONSucceeds(t, `{"meter":"4/4", "partNotes":{}}`)
// }

// func testValidateJSONSucceeds(t *testing.T, jsonStr string) {
// 	loader := gojsonschema.NewStringLoader(jsonStr)

// 	assert.NoError(t, measure.ValidateJSON(loader))
// }

// func testValidateJSONFails(t *testing.T, jsonStr string) {
// 	loader := gojsonschema.NewStringLoader(jsonStr)

// 	assert.Error(t, measure.ValidateJSON(loader))
// }
