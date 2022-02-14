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

func TestScoreProgramSections(t *testing.T) {
	s := score.New()

	secs := s.ProgramSections()

	assert.Empty(t, secs)

	s.Program = append(s.Program, "section1")
	secs = s.ProgramSections()

	assert.Empty(t, secs)

	s.Sections["section1"] = section.New()
	secs = s.ProgramSections()

	assert.Len(t, secs, 1)

	s.Program = append(s.Program, "section1")
	secs = s.ProgramSections()

	assert.Len(t, secs, 2)

	s.Sections["section2"] = section.New()
	s.Program = append(s.Program, "section2")
	secs = s.ProgramSections()

	assert.Len(t, secs, 3)
}
