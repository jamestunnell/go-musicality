package section_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/section"
)

func TestSectionEmpty(t *testing.T) {
	s := section.New()

	assert.Nil(t, s.Validate())
}

func TestSectionNotEmpty(t *testing.T) {
	s := section.New()

	s.AppendMeasures(1, meter.New(4, 4))

	assert.Nil(t, s.Validate())
}

func TestSectionInvalid(t *testing.T) {
	s := section.New()

	s.AppendMeasures(1, meter.New(0, 4))

	assert.NotNil(t, s.Validate())
}
