package score_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

func TestValidateJSON(t *testing.T) {
	validJSON := `
	{
		"program": ["section1"],
		"sections": {
			"section1": {
				"startTempo": 120,
				"startDynamic": 0.0,
				"measures": [
					{
						"meter": "4/4",
						"partNotes": {
							"piano": [
								{"duration": "3/4" },
								{"duration": "1/4", "pitches": ["C4"]}
							]
						}
					}
				]
			}
		}
	}`

	testValidateJSONValid(t, "happy path", validJSON)

	testValidateJSONInvalid(t, "missing program propert", strings.Replace(validJSON, "program", "not-program", 1))

	testValidateJSONInvalid(t, "missing sections property", strings.Replace(validJSON, "sections", "not-sections", 1))

	testValidateJSONInvalid(t, "program is wrong type", `
	{
		"program": {"a": 1},
		"sections": {}
	}`)

	testValidateJSONInvalid(t, "sections is wrong type", `
	{
		"program": ["section1"],
		"sections": [1,2,3]
	}`)

	testValidateJSONInvalid(t, "bad meter", strings.Replace(validJSON, `"meter": "4/4"`, `"meter": "1/0"`, 1))

	testValidateJSONInvalid(t, "meter wrong type", strings.Replace(validJSON, `"meter": "4/4"`, `"meter": 5`, 1))

	testValidateJSONInvalid(t, "bad pitch string", strings.Replace(validJSON, "C4", "H2", 1))
}

func testValidateJSONValid(t *testing.T, name, jsonStr string) {
	t.Run(name, func(t *testing.T) {
		loader := gojsonschema.NewStringLoader(jsonStr)

		result, err := score.Schema().Validate(loader)

		require.NoError(t, err)
		assert.True(t, result.Valid())
		assert.Empty(t, result.Errors())
	})
}

func testValidateJSONInvalid(t *testing.T, name, jsonStr string) {
	t.Run(name, func(t *testing.T) {
		loader := gojsonschema.NewStringLoader(jsonStr)

		result, err := score.Schema().Validate(loader)

		require.NoError(t, err)
		assert.False(t, result.Valid())
		assert.NotEmpty(t, result.Errors())
	})
}
