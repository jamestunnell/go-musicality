package score_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"

	"github.com/jamestunnell/go-musicality/common/value"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/score"
)

func TestValidateJSON(t *testing.T) {
	makeValidScoreMap := func() value.Map {
		s := value.Map{
			"program": value.Slice{"section1"},
			"sections": value.Map{
				"section1": value.Map{
					"startTempo":   120,
					"startDynamic": 0.0,
					"startMeter": value.Map{
						"beatsPerMeasure": 4,
						"beatDuration":    "1/4",
					},
					"measures": value.Slice{
						value.Map{
							"partNotes": value.Map{
								"piano": value.Slice{
									value.Map{"duration": "3/4"},
									value.Map{
										"duration":   "1/4",
										"pitches":    value.Slice{"C4"},
										"attack":     note.ControlMax,
										"separation": note.ControlMax,
										"links": value.Map{
											"C4": value.Map{
												"target": "C4",
												"type":   "tie",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		return s
	}

	// Valid scores
	testValidateJSONValid(t, "happy path", makeValidScoreMap(), func(m value.Map) {})
	testValidateJSONValid(t, "measure with no part notes", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("partNotes", value.Map{}))
	})
	testValidateJSONValid(t, "note with no attack", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("attack"))
	})
	testValidateJSONValid(t, "note with no separation", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("separation"))
	})
	testValidateJSONValid(t, "measure with dynamic change", makeValidScoreMap(), func(m value.Map) {
		val, found := m.GetSliceValue("measures", 0)

		require.True(t, found)

		mm, ok := val.(value.Map)

		require.True(t, ok)

		mm["dynamicChanges"] = value.Map{
			"1/4": value.Map{
				"endValue": 1.0,
				"duration": "3/4",
			},
		}
	})

	// Invalid scores
	testValidateJSONInvalid(t, "missing program", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("program"))
	})
	testValidateJSONInvalid(t, "missing sections", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("sections"))
	})
	testValidateJSONInvalid(t, "program wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("program", value.Map{}))
	})
	testValidateJSONInvalid(t, "sections wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("sections", value.Slice{}))
	})
	testValidateJSONInvalid(t, "empty measure", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("measures", value.Slice{value.Map{}}))
	})
	testValidateJSONInvalid(t, "missing part notes", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("partNotes"))
	})
	testValidateJSONInvalid(t, "bad meter", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("beatDuration", 42.5))
	})
	testValidateJSONInvalid(t, "meter wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("startMeter", 5.5))
	})
	testValidateJSONInvalid(t, "bad pitch", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("pitches", value.Slice{"not-a-pitch"}))
	})
	testValidateJSONInvalid(t, "bad duration", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("duration", 2.5))
	})
	testValidateJSONInvalid(t, "duration wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("duration", 5.6))
	})
	testValidateJSONInvalid(t, "missing duration", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.RemoveValue("duration"))
	})
	testValidateJSONInvalid(t, "bad attack", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("attack", note.ControlMin-0.01))
	})
	testValidateJSONInvalid(t, "bad attack", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("attack", note.ControlMax+0.01))
	})
	testValidateJSONInvalid(t, "bad separation", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("separation", note.ControlMin-0.01))
	})
	testValidateJSONInvalid(t, "bad separation", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("separation", note.ControlMax+0.01))
	})
	testValidateJSONInvalid(t, "bad link target pitch", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("target", "not-a-pitch"))
	})
	testValidateJSONInvalid(t, "link target pitch wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("target", 10.5))
	})
	testValidateJSONInvalid(t, "unknown link type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("type", "bad-link"))
	})
	testValidateJSONInvalid(t, "link type wrong type", makeValidScoreMap(), func(m value.Map) {
		assert.True(t, m.ChangeValue("type", 22.2))
	})
}

func testValidateJSONValid(
	t *testing.T,
	name string,
	scoreMap value.Map,
	mod value.MapModFunc) {
	t.Run(name, func(t *testing.T) {
		result, err := testValidateJSON(t, name, scoreMap, mod)

		require.NoError(t, err)
		assert.True(t, result.Valid())
		assert.Empty(t, result.Errors())
	})
}

func testValidateJSONInvalid(
	t *testing.T,
	name string,
	scoreMap value.Map,
	mod value.MapModFunc) {
	t.Run(name, func(t *testing.T) {
		result, err := testValidateJSON(t, name, scoreMap, mod)

		require.NoError(t, err)
		assert.False(t, result.Valid())
		assert.NotEmpty(t, result.Errors())
	})
}

func testValidateJSON(
	t *testing.T,
	name string,
	scoreMap value.Map,
	mod value.MapModFunc) (*gojsonschema.Result, error) {
	mod(scoreMap)

	jsonData, err := json.Marshal(scoreMap)

	require.NoError(t, err)

	loader := gojsonschema.NewStringLoader(string(jsonData))

	return score.Schema().Validate(loader)
}
