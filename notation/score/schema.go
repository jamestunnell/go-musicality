package score

import "github.com/xeipuuv/gojsonschema"

var schemaStr = `
{
	"$schema": "http://json-schema.org/draft-07/schema",
	"$id": "http://github.com/jamestunnell/go-musicality/score.json",
	"type": "object",
	"title": "Score JSON Schema",
	"description": "JSON schema for a music score.",
	"examples": [
		{}
	],
	"required": [
		"program",
		"sections"
	],
	"properties": {
		"sections": {
			"$id": "#/properties/sections",
			"type": "object",
			"title": "Score sections",
			"description": "Organizes score into logical sections.",
			"additionalProperties": {
				"$ref": "#/definitions/section"
			}
		},
		"program": {
			"$id": "#/properties/program",
			"type": "array",
			"title": "Score program",
			"description": "Determines playback of score sections.",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"settings": {
			"$id": "#/properties/settings",
			"title": "Score settings",
			"description": "Holds format-specific settings, split into categories.",
			"type": "object"
		}
	},
	"definitions": {
		"section": {
			"$id": "#/definitions/section",
			"title": "Score section",
			"description": "Contains a logical piece of the score",
			"required": [
				"startTempo",
				"startDynamic",
				"measures"
			],
			"properties": {
				"startTempo": {
					"$ref": "#/definitions/tempo"
				},
				"startDynamic": {
					"$ref": "#/definitions/dynamic"
				},
				"measures": {
					"type": "array",
					"title": "Section measures",
					"description": "Contains section part notes.",
					"items": {
						"$ref": "#/definitions/measure"
					}
				}
			}
		},
		"tempo": {
			"$id": "#/definitions/tempo",
			"type": "number",
			"title": "Tempo",
			"description": "Tempo (beats per minute).",
			"exclusiveMinimum": 0
		},
		"dynamic": {
			"$id": "#/definitions/dynamic",
			"title": "Dynamic",
			"description": "Controls loudness.",
			"$ref": "#/definitions/control"
		},
		"measure": {
			"$id": "#/definitions/measure",
			"type": "object",
			"title": "Measure",
			"description": "One measure with notes divided by part.",
			"required": [
				"meter",
				"partNotes"
			],
			"properties": {
				"meter": {
					"$ref": "#/definitions/meter"
				},
				"partNotes": {
					"$id": "#/properties/partNotes",
					"type": "object",
					"title": "Part notes",
					"description": "Maps part names to measure notes.",
					"default": {},
					"additionalProperties": {
						"type": "array",
						"minItems": 1,
						"items": {
							"$ref": "#/definitions/note"
						}
					}
				},
				"dynamicChanges": {
					"type": "array",
					"title": "Dynamic changes",
					"items": {
						"$ref": "#/definitions/change"
					}
				},
				"tempoChanges": {
					"type": "array",
					"title": "Tempo changes",
					"items": {
						"$ref": "#/definitions/change"
					}
				}
			}
		},
		"change": {
			"$id": "#/definitions/change",
			"title": "A value change",
			"type": "object",
			"properties": {
				"offset": {"$ref": "#/definitions/rationalNonNegative"},
				"endValue": {"type": "number"},
				"duration": {"$ref": "#/definitions/duration"},
			}
		},
		"meter": {
			"$id": "#/definitions/meter",
			"type": "string",
			"title": "Meter",
			"description": "Measure time signature",
			"pattern": "^[1-9][0-9]*/[1-9][0-9]*$",
			"examples": [
				"4/4"
			]
		},
		"note": {
			"type": "object",
			"title": "Note",
			"description": "Note with duration and zero or more pithches.",
			"required": [
				"duration"
			],
			"properties": {
				"duration": {
					"$ref": "#/definitions/duration"
				},
				"separation": {
					"$ref": "#/definitions/control",
					"$id": "#/properties/separation",
					"title": "Note separation",
					"description": "Separation between current and next note."
				},
				"attack": {
					"$id": "#/properties/attack",
					"$ref": "#/definitions/control",
					"title": "Note attack",
					"description": "Emphasis at note onset."
				},
				"pitches": {
					"$ref": "#/definitions/pitches"
				},
				"links": {
					"$id": "#/properties/links",
					"type": "object",
					"title": "Links",
					"description": "Links from pitches in the current note to pitches in the next note.",
					"default": {},
					"patternProperties": {
						"^[A-G][#b]?-?[0-9]+$": {
							"$ref": "#/definitions/link"
						}
					},
					"additionalProperties": false
				}
			}
		},
		"control": {
			"$id": "#/definitions/control",
			"type": "number",
			"minimum": -1,
			"maximum": 1,
			"default": 0
		},
		"rationalNonNegative": {
			"$id": "#/definitions/rationalNonNegative",
			"type": "string",
			"title": "Rational number",
			"pattern": "^[0-9]+/[1-9][0-9]*$",
			"examples": [
				"0/1",
				"1/2",
				"1/1",
			]
		}
		"duration": {
			"$id": "#/definitions/duration",
			"type": "string",
			"title": "Note duration",
			"description": "Nominal note length.",
			"pattern": "^[1-9][0-9]*/[1-9][0-9]*$",
			"examples": [
				"1/2"
			]
		},
		"pitches": {
			"$id": "#/definitions/pitches",
			"type": "array",
			"title": "Pitches",
			"description": "Set of note pitches, represented as strings.",
			"default": [],
			"minItems": 0,
			"items": {
				"type": "string",
				"pattern": "^[A-G][#b]?-?[0-9]+$"
			},
			"examples": [
				[
					"C4",
					"E4",
					"Bb2"
				]
			]
		},
		"link": {
			"$id": "#/definitions/link",
			"type": "object",
			"title": "Link",
			"description": "Link from pitch in the current note to a pitch in the next note.",
			"examples": [
				{
					"type": "tie",
					"target": "C4"
				}
			],
			"required": [
				"type",
				"target"
			],
			"properties": {
				"type": {
					"$ref": "#/definitions/linkType"
				},
				"target": {
					"$ref": "#/definitions/linkTarget"
				}
			},
			"additionalProperties": true
		},
		"linkType": {
			"$id": "#/definitions/linkType",
			"type": "string",
			"title": "Link type",
			"enum": [
				"tie",
				"glide",
				"step"
			],
			"examples": [
				"tie"
			]
		},
		"linkTarget": {
			"$id": "#/definitions/linkTarget",
			"type": "string",
			"title": "Link target",
			"pattern": "^[A-G][#b]?-?[0-9]+$",
			"description": "Pitch in the next note that will be linked to.",
			"default": "",
			"examples": [
				"C4"
			]
		}
	}
}`

var schema *gojsonschema.Schema

func init() {
	var err error

	schema, err = gojsonschema.NewSchema(SchemaLoader())

	if err != nil {
		panic(err)
	}
}

func SchemaLoader() gojsonschema.JSONLoader {
	return gojsonschema.NewStringLoader(schemaStr)
}

func Schema() *gojsonschema.Schema {
	return schema
}
