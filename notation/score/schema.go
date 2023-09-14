package score

import "github.com/xeipuuv/gojsonschema"

const schemaStr = `
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
				"startMeter",
				"measures"
			],
			"properties": {
				"startTempo": {
					"$ref": "#/definitions/tempo"
				},
				"startMeter": {
					"$ref": "#/definitions/meter"
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
			"$ref": "#/definitions/control",
			"title": "Dynamic",
			"description": "Controls loudness."
		},
		"measure": {
			"$id": "#/definitions/measure",
			"type": "object",
			"title": "Measure",
			"description": "One measure with notes divided by part.",
			"required": [
				"partNotes"
			],
			"properties": {
				"meterChange": {
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
					"title": "Dynamic changes",
					"$ref": "#/definitions/changeMap"
				},
				"tempoChanges": {
					"title": "Tempo changes",
					"$ref": "#/definitions/changeMap"
				}
			}
		},
		"changes": {
			"$id": "#/definitions/changeMap",
			"type": "object",
			"title": "Change map",
			"description": "Maps changes by offset",
			"patternProperties": {
				"^-?[0-9]+(/[1-9][0-9]*)?$": {
					"$ref": "#/definitions/change"
				}
			},
			"additionalProperties": false
		},
		"change": {
			"$id": "#/definitions/change",
			"title": "A value change",
			"type": "object",
			"properties": {
				"offset": { "$ref": "#/definitions/rational"},
				"endValue": {"type": "number"},
				"duration": {"$ref": "#/definitions/rational"}
			}
		},
		"meter": {
			"$id": "#/definitions/meter",
			"type": "object",
			"title": "Meter",
			"description": "Measure time signature",
			"required": ["beatDuration","beatsPerMeasure"],
			"properties": {
				"beatDuration": {"$ref": "#/definitions/rational"},
				"beatsPerMeasure": {
					"type": "integer",
					"exclusiveMinimum": 0
				}
			},
			"examples": [
				{"beatDuration":"1/4","beatsPerMeasure":4}
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
					"$ref": "#/definitions/rational"
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
			"minimum": 0,
			"maximum": 1,
			"default": 0.5
		},
		"rational": {
			"$id": "#/definitions/rational",
			"type": "string",
			"title": "Rational number",
			"description": "Numerator and optional denominator.",
			"pattern": "^-?[0-9]+(/[1-9][0-9]*)?$",
			"examples": [
				"-2",
				"-1/2",
				"7",
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
				"step",
				"slur",
				"stepSlurred"
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
