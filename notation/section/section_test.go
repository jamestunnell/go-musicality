package section_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/measure"
	"github.com/jamestunnell/go-musicality/notation/meter"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/pitch"
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

	s = section.New(section.OptStartMeter(meter.FourFour()))

	assert.Nil(t, s.Validate())
	assert.True(t, s.StartMeter.Equal(meter.FourFour()))

	s = section.New(section.OptStartMeter(meter.New(0, rat.New(1, 4))))

	assert.NotNil(t, s.Validate())
}

func TestSectionEmpty(t *testing.T) {
	s := section.New()

	assert.Nil(t, s.Validate())
}

func TestSectionNotEmpty(t *testing.T) {
	s := section.New()
	m := measure.New()

	s.Measures = append(s.Measures, m)

	assert.Nil(t, s.Validate())
}

func TestSectionInvalid(t *testing.T) {
	s := section.New()
	m := measure.New()

	m.MeterChange = meter.New(0, rat.New(1, 4))

	s.Measures = append(s.Measures, m)

	assert.NotNil(t, s.Validate())
}

func TestSectionDuration(t *testing.T) {
	s := section.New()

	assert.True(t, s.Duration().Zero())

	s.StartMeter = meter.ThreeFour()

	m := measure.New()

	s.Measures = append(s.Measures, m)

	assert.True(t, s.Duration().Equal(rat.New(3, 4)))

	m2 := measure.New()

	m2.MeterChange = meter.TwoFour()

	s.Measures = append(s.Measures, m2)

	assert.True(t, s.Duration().Equal(rat.New(5, 4)))
}

func TestSectionParts(t *testing.T) {
	s := section.New()

	s.StartMeter = meter.ThreeFour()

	m1 := measure.New()
	m2 := measure.New()
	n1 := note.Quarter(pitch.C4)
	n2 := note.Half(pitch.C2)

	m2.MeterChange = meter.FourFour()

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

func TestDynamicChanges(t *testing.T) {
	s := section.New(
		section.OptStartDynamic(1.0),
		section.OptStartMeter(meter.FourFour()),
	)
	m1 := measure.New()
	m2 := measure.New()
	m3 := measure.New()

	s.Measures = append(s.Measures, m1, m2, m3)

	dc := s.DynamicChanges(rat.Zero())

	assert.Empty(t, dc)

	m2.DynamicChanges = append(m2.DynamicChanges, change.NewImmediate(rat.Zero(), 0.8))

	dc = s.DynamicChanges(rat.Zero())

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(1, 1)))

	dc = s.DynamicChanges(rat.New(4, 1))

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(5, 1)))
}

func TestTempoChanges(t *testing.T) {
	s := section.New(
		section.OptStartTempo(100),
		section.OptStartMeter(meter.FourFour()),
	)
	m1 := measure.New()
	m2 := measure.New()
	m3 := measure.New()

	s.Measures = append(s.Measures, m1, m2, m3)

	dc := s.TempoChanges(rat.Zero())

	assert.Empty(t, dc)

	m2.TempoChanges = append(m2.TempoChanges, change.NewImmediate(rat.Zero(), 90))

	dc = s.TempoChanges(rat.Zero())

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(1, 1)))

	dc = s.TempoChanges(rat.New(4, 1))

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(5, 1)))
}

func TestBeatDurChanges(t *testing.T) {
	s := section.New(
		section.OptStartMeter(meter.FourFour()),
	)
	m1 := measure.New()
	m2 := measure.New()
	m3 := measure.New()

	s.Measures = append(s.Measures, m1, m2, m3)

	dc := s.BeatDurChanges(rat.Zero())

	assert.Empty(t, dc)

	m2.MeterChange = meter.ThreeFour()

	dc = s.BeatDurChanges(rat.Zero())

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(1, 1)))

	dc = s.BeatDurChanges(rat.New(4, 1))

	require.Len(t, dc, 1)

	assert.True(t, dc[0].Offset.Equal(rat.New(5, 1)))
}
