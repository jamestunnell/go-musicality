package section_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/section"
)

func TestNewWithOpt(t *testing.T) {
	s := section.New(section.OptStartDynamic(0.75))

	assert.Nil(t, s.Validate())
	assert.Equal(t, 0.75, s.StartDynamic)

	s = section.New(section.OptStartDynamic(1.75))

	assert.NotNil(t, s.Validate())

	s = section.New(section.OptStartTempo(200))

	assert.Nil(t, s.Validate())
	assert.Equal(t, float64(200), s.StartTempo)

	s = section.New(section.OptStartTempo(0))

	assert.NotNil(t, s.Validate())
}

func TestSectionEmpty(t *testing.T) {
	s := section.New()

	assert.Nil(t, s.Validate())
}

func TestSectionNotEmpty(t *testing.T) {
	s := section.New()
	m := measure.New(meter.New(4, 4))

	s.Measures = append(s.Measures, m)

	assert.Nil(t, s.Validate())
}

func TestSectionInvalid(t *testing.T) {
	s := section.New()
	m := measure.New(meter.New(0, 4))

	s.Measures = append(s.Measures, m)

	assert.NotNil(t, s.Validate())
}

func TestSectionDuration(t *testing.T) {
	s := section.New()

	assert.True(t, s.Duration().Zero())

	m := measure.New(meter.New(3, 4))

	s.Measures = append(s.Measures, m)

	assert.True(t, s.Duration().Equal(rat.New(3, 4)))

	m2 := measure.New(meter.New(2, 4))

	s.Measures = append(s.Measures, m2)

	assert.True(t, s.Duration().Equal(rat.New(5, 4)))
}

func TestSectionParts(t *testing.T) {
	s := section.New()
	m1 := measure.New(meter.New(3, 4))
	m2 := measure.New(meter.New(4, 4))
	n1 := note.Quarter(pitch.C4)
	n2 := note.Half(pitch.C2)

	m1.PartNotes["piano"] = []*note.Note{note.Half(), n1}
	m2.PartNotes["bass"] = []*note.Note{note.Half(), n2}

	s.Measures = append(s.Measures, m1, m2)

	require.ElementsMatch(t, []string{"piano", "bass"}, s.PartNames())

	pianoNotes := s.PartNotes("piano")
	bassNotes := s.PartNotes("bass")

	require.Len(t, pianoNotes, 3)
	require.Len(t, bassNotes, 3)

	assert.True(t, pianoNotes[0].Equal(note.Half()))
	assert.True(t, pianoNotes[1].Equal(n1))
	assert.True(t, pianoNotes[2].Equal(note.Whole()))

	assert.True(t, bassNotes[0].Equal(note.New(rat.New(3, 4))))
	assert.True(t, bassNotes[1].Equal(note.Half()))
	assert.True(t, bassNotes[2].Equal(n2))
}

// func TestDynamics(t *testing.T) {
// 	s := section.New()
// 	m1 := measure.New(meter.New(4, 4))
// 	m2 := measure.New(meter.New(4, 4))
// 	ch1 := change.New(0.0, rat.New(1, 1))
// 	ch2 := change.NewImmediate(1.0)

// 	assert.Empty(t, s.DynamicChanges(rat.Zero()))

// 	m1.DynamicChanges[rat.Zero()] = ch1
// 	m2.DynamicChanges[rat.New(1, 2)] = ch2

// 	s.Measures = append(s.Measures, m1, m2)

// 	dc := s.DynamicChanges(rat.Zero())
// 	off1 := rat.Zero()

// 	require.Len(t, dc, 2)

// 	c := dc[off1]

// 	require.NotNil(t, c)
// 	// require.NotNil(t, dc[off2])

// 	// assert.True(t, dc[off1].Equal(ch1))
// 	// assert.True(t, dc[off2].Equal(ch2))
// }
