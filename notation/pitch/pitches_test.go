package pitch_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func TestPitchesSort(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C3, pitch.C2}

	sort.Sort(ps)

	assert.True(t, ps[0].Equal(pitch.C1))
	assert.True(t, ps[1].Equal(pitch.C2))
	assert.True(t, ps[2].Equal(pitch.C3))
}

func TestPitchesCombinationFrom0Pitches(t *testing.T) {
	ps := pitch.Pitches{}

	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{})
}

func TestPitchesCombinationFrom1Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1}

	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
	})
}

func TestPitchesCombinationFrom2Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2}

	testComb(t, ps, -1, []pitch.Pitches{})
	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
	})
	testComb(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
	})
	testComb(t, ps, 3, []pitch.Pitches{})
}

func TestPitchesCombinationFrom3Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2, pitch.C3}

	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
		{pitch.C3},
	})
	testComb(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
		{pitch.C1, pitch.C3},
		{pitch.C2, pitch.C3},
	})
	testComb(t, ps, 3, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3},
	})
}

func TestPitchesCombinationFrom4Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2, pitch.C3, pitch.C4}

	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
		{pitch.C3},
		{pitch.C4},
	})
	testComb(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
		{pitch.C1, pitch.C3},
		{pitch.C1, pitch.C4},
		{pitch.C2, pitch.C3},
		{pitch.C2, pitch.C4},
		{pitch.C3, pitch.C4},
	})
	testComb(t, ps, 3, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3},
		{pitch.C1, pitch.C2, pitch.C4},
		{pitch.C1, pitch.C3, pitch.C4},
		{pitch.C2, pitch.C3, pitch.C4},
	})
	testComb(t, ps, 4, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3, pitch.C4},
	})
}

func testComb(t *testing.T, ps pitch.Pitches, n int, expectedCombs []pitch.Pitches) {
	name := fmt.Sprintf("comb %d from %d pitches", n, ps.Len())

	t.Run(name, func(t *testing.T) {
		combs := []pitch.Pitches{}
		gatherCombs := func(comb pitch.Pitches) {
			sort.Sort(comb)
			combs = append(combs, comb)
		}

		ps.Combination(n, gatherCombs)

		assert.ElementsMatch(t, expectedCombs, combs)
	})
}
