package score_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/score"
)

func TestScoreValidSection(t *testing.T) {
	s := &score.Score{
		Start: &score.State{
			Tempo:  120.0,
			Volume: 0.5,
		},
		Sections: []*score.Section{
			{Name: "", Measures: []*measure.Measure{measure.New(meter.New(4, 4))}},
		},
	}

	assert.Nil(t, s.Validate())
}

func TestScoreInvalidStart(t *testing.T) {
	s := &score.Score{
		Start: &score.State{
			Tempo:  0.0,
			Volume: 0.5,
		},
		Sections: []*score.Section{},
	}

	assert.NotNil(t, s.Validate())
}

func TestScoreInvalidSection(t *testing.T) {
	s := &score.Score{
		Start: &score.State{
			Tempo:  120.0,
			Volume: 0.5,
		},
		Sections: []*score.Section{
			{Name: "", Measures: []*measure.Measure{measure.New(meter.New(0, 4))}},
		},
	}

	assert.NotNil(t, s.Validate())
}
