package score_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
)

func TestScoreValidSection(t *testing.T) {
	s := score.New()

	s.Program = append(s.Program, "notempty")
	s.Sections["notempty"] = section.New(
		section.OptStartTempo(120),
		section.OptStartDynamic(0.0),
	)

	assert.Nil(t, s.Validate())
}

func TestScoreInvalidSection(t *testing.T) {
	sec := section.New(section.OptStartTempo(0.0))
	s := score.New()

	s.Sections["notempty"] = sec

	assert.NotNil(t, s.Validate())
}

func TestScoreMissingSection(t *testing.T) {
	s := score.New()

	s.Program = append(s.Program, "notempty")

	assert.NotNil(t, s.Validate())
}
