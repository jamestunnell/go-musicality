package score_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
)

func TestScoreValidSection(t *testing.T) {
	s := &score.Score{
		Start: &score.State{
			Tempo:  120.0,
			Volume: 0.5,
		},
		Sections: []*section.Section{
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
		Sections: []*section.Section{},
	}

	assert.NotNil(t, s.Validate())
}

func TestScoreInvalidSection(t *testing.T) {
	sec := section.New("")

	sec.AppendMeasures(1, meter.New(0, 4))

	s := &score.Score{
		Start: &score.State{
			Tempo:  120.0,
			Volume: 0.5,
		},
		Sections: []*section.Section{sec},
	}

	assert.NotNil(t, s.Validate())
}
