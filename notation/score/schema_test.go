package score_test

import (
	"encoding/json"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

type Value = interface{}
type Map = map[string]Value
type Slice = []Value

type MapModFunc func(m Map)

func TestValidateJSON(t *testing.T) {
	makeValidScoreMap := func() Map {
		s := Map{
			"program": Slice{"section1"},
			"sections": Map{
				"section1": Map{
					"startTempo":   120,
					"startDynamic": 0.0,
					"measures": Slice{
						Map{
							"meter": "4/4",
							"partNotes": Map{
								"piano": Slice{
									Map{"duration": "3/4"},
									Map{
										"duration":   "1/4",
										"pitches":    Slice{"C4"},
										"attack":     0.0,
										"separation": 0.0,
										"links": Map{
											"C4": Map{
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
	testValidateJSONValid(t, "happy path", makeValidScoreMap(), func(m Map) {})
	testValidateJSONValid(t, "measure with no part notes", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "partNotes", Map{}))
	})

	// Invalid scores
	testValidateJSONInvalid(t, "missing program", makeValidScoreMap(), func(m Map) {
		assert.True(t, RemoveMapValue(m, "program"))
	})
	testValidateJSONInvalid(t, "missing sections", makeValidScoreMap(), func(m Map) {
		assert.True(t, RemoveMapValue(m, "sections"))
	})
	testValidateJSONInvalid(t, "program wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "program", Map{}))
	})
	testValidateJSONInvalid(t, "sections wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "sections", Slice{}))
	})
	testValidateJSONInvalid(t, "empty measure", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "measures", Slice{Map{}}))
	})
	testValidateJSONInvalid(t, "missing part notes", makeValidScoreMap(), func(m Map) {
		assert.True(t, RemoveMapValue(m, "partNotes"))
	})
	testValidateJSONInvalid(t, "bad meter", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "meter", "0/4"))
	})
	testValidateJSONInvalid(t, "meter wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "meter", 5.5))
	})
	testValidateJSONInvalid(t, "bad pitch", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "pitches", Slice{"not-a-pitch"}))
	})
	testValidateJSONInvalid(t, "bad duration", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "duration", "0/1"))
	})
	testValidateJSONInvalid(t, "duration wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "duration", 5.6))
	})
	testValidateJSONInvalid(t, "missing duration", makeValidScoreMap(), func(m Map) {
		assert.True(t, RemoveMapValue(m, "duration"))
	})
	testValidateJSONInvalid(t, "bad attack", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "attack", -1.01))
	})
	testValidateJSONInvalid(t, "bad attack", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "attack", 1.01))
	})
	testValidateJSONInvalid(t, "bad separation", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "separation", -1.01))
	})
	testValidateJSONInvalid(t, "bad separation", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "separation", 1.01))
	})
	testValidateJSONInvalid(t, "bad link target pitch", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "target", "not-a-pitch"))
	})
	testValidateJSONInvalid(t, "link target pitch wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "target", 10.5))
	})
	testValidateJSONInvalid(t, "unknown link type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "type", "bad-link"))
	})
	testValidateJSONInvalid(t, "link type wrong type", makeValidScoreMap(), func(m Map) {
		assert.True(t, ChangeMapValue(m, "type", 22.2))
	})
}

func testValidateJSONValid(t *testing.T, name string, scoreMap Map, mod MapModFunc) {
	t.Run(name, func(t *testing.T) {
		result, err := testValidateJSON(t, name, scoreMap, mod)

		require.NoError(t, err)
		assert.True(t, result.Valid())
		assert.Empty(t, result.Errors())
	})
}

func testValidateJSONInvalid(t *testing.T, name string, scoreMap Map, mod MapModFunc) {
	t.Run(name, func(t *testing.T) {
		result, err := testValidateJSON(t, name, scoreMap, mod)

		require.NoError(t, err)
		assert.False(t, result.Valid())
		assert.NotEmpty(t, result.Errors())
	})
}

func testValidateJSON(t *testing.T, name string, scoreMap Map, mod MapModFunc) (*gojsonschema.Result, error) {
	mod(scoreMap)

	jsonData, err := json.Marshal(scoreMap)

	require.NoError(t, err)

	loader := gojsonschema.NewStringLoader(string(jsonData))

	return score.Schema().Validate(loader)
}

type KeyMatchActionFunc func(m Map, k string, v Value) bool

// ExecuteOnKeyMatch searches the entire container recursively looking for the
// first key matching the given key, and executes the given action function.
// Recursive searching and action execution continues until the action function
// return true.
func ExecuteOnKeyMatch(v Value, keyMatch string, f KeyMatchActionFunc) bool {
	switch vv := v.(type) {
	case Map:
		for k, vvv := range vv {
			if k == keyMatch {
				if f(vv, k, vvv) {
					return true
				}
			}

			if ExecuteOnKeyMatch(vvv, keyMatch, f) {
				return true
			}
		}

	case Slice:
		for _, vvv := range vv {
			if ExecuteOnKeyMatch(vvv, keyMatch, f) {
				return true
			}
		}
	}

	return false
}

// ChangeMapValue searches the entire container recursively looking for the
// first key matching the given key, and changes the associated value in the parent map.
func ChangeMapValue(rootMap Map, matchKey string, newValue Value) bool {
	f := func(parentMap Map, k string, v Value) bool {
		parentMap[k] = newValue

		return true
	}

	return ExecuteOnKeyMatch(rootMap, matchKey, f)
}

// RemoveMapKey searches the entire container recursively looking for the
// first key matching the given key, and removes the key from the parent map.
func RemoveMapValue(rootMap Map, matchKey string) bool {
	f := func(parentMap Map, k string, v Value) bool {
		delete(parentMap, matchKey)

		return true
	}

	return ExecuteOnKeyMatch(rootMap, matchKey, f)
}

func ChangeSliceValue(rootMap Map, sliceKey string, index int, newValue Value) bool {
	f := func(parentMap Map, k string, v Value) bool {
		s, ok := v.(Slice)
		if !ok {
			return false
		}

		s[index] = newValue
		// parentMap[k] = newValue

		return true
	}

	return ExecuteOnKeyMatch(rootMap, sliceKey, f)
}
