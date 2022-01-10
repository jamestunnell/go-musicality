package score_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/score"
)

func TestSectionEmpty(t *testing.T) {
	s := &score.Section{
		Name:     "",
		Measures: []*measure.Measure{},
	}

	assert.Nil(t, s.Validate())
}

func TestSectionNotEmpty(t *testing.T) {
	s := &score.Section{
		Name: "anything",
		Measures: []*measure.Measure{
			measure.New(meter.New(4, 4)),
		},
	}

	assert.Nil(t, s.Validate())
}

func TestSectionInvalid(t *testing.T) {
	s := &score.Section{
		Name: "anything",
		Measures: []*measure.Measure{
			measure.New(meter.New(0, 4)),
		},
	}

	assert.NotNil(t, s.Validate())
}
