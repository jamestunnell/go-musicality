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

func TestPitchesCombAndPermFrom0Pitches(t *testing.T) {
	ps := pitch.Pitches{}

	// Combinations
	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{})

	// Permutations
	testPerm(t, ps, 0, []pitch.Pitches{
		{},
	})
	testPerm(t, ps, 1, []pitch.Pitches{})
}

func TestPitchesCombAndPermFrom1Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1}

	// Combinations
	testComb(t, ps, 0, []pitch.Pitches{
		{},
	})
	testComb(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
	})

	// Permutations
	testPerm(t, ps, 0, []pitch.Pitches{
		{},
	})
	testPerm(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
	})
}

func TestPitchesCombAndPermFrom2Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2}

	// Combinations
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

	// Permutations
	testPerm(t, ps, -1, []pitch.Pitches{})
	testPerm(t, ps, 0, []pitch.Pitches{
		{},
	})
	testPerm(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
	})
	testPerm(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
		{pitch.C2, pitch.C1},
	})
	testPerm(t, ps, 3, []pitch.Pitches{})
}

func TestPitchesCombAndPermFrom3Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2, pitch.C3}

	// Combinations
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

	// Permutations
	testPerm(t, ps, 0, []pitch.Pitches{
		{},
	})
	testPerm(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
		{pitch.C3},
	})
	testPerm(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
		{pitch.C1, pitch.C3},
		{pitch.C2, pitch.C1},
		{pitch.C2, pitch.C3},
		{pitch.C3, pitch.C1},
		{pitch.C3, pitch.C2},
	})
	testPerm(t, ps, 3, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3},
		{pitch.C1, pitch.C3, pitch.C2},
		{pitch.C2, pitch.C1, pitch.C3},
		{pitch.C2, pitch.C3, pitch.C1},
		{pitch.C3, pitch.C1, pitch.C2},
		{pitch.C3, pitch.C2, pitch.C1},
	})
	testPerm(t, ps, 4, []pitch.Pitches{})
}

func TestPitchesCombAndPermFrom4Pitches(t *testing.T) {
	ps := pitch.Pitches{pitch.C1, pitch.C2, pitch.C3, pitch.C4}

	// Combinations
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

	// Permutations
	testPerm(t, ps, 0, []pitch.Pitches{
		{},
	})
	testPerm(t, ps, 1, []pitch.Pitches{
		{pitch.C1},
		{pitch.C2},
		{pitch.C3},
		{pitch.C4},
	})
	testPerm(t, ps, 2, []pitch.Pitches{
		{pitch.C1, pitch.C2},
		{pitch.C1, pitch.C3},
		{pitch.C1, pitch.C4},
		{pitch.C2, pitch.C1},
		{pitch.C2, pitch.C3},
		{pitch.C2, pitch.C4},
		{pitch.C3, pitch.C1},
		{pitch.C3, pitch.C2},
		{pitch.C3, pitch.C4},
		{pitch.C4, pitch.C1},
		{pitch.C4, pitch.C2},
		{pitch.C4, pitch.C3},
	})
	testPerm(t, ps, 3, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3},
		{pitch.C1, pitch.C2, pitch.C4},
		{pitch.C1, pitch.C3, pitch.C4},
		{pitch.C1, pitch.C3, pitch.C2},
		{pitch.C1, pitch.C4, pitch.C2},
		{pitch.C1, pitch.C4, pitch.C3},

		{pitch.C2, pitch.C1, pitch.C3},
		{pitch.C2, pitch.C1, pitch.C4},
		{pitch.C2, pitch.C3, pitch.C1},
		{pitch.C2, pitch.C3, pitch.C4},
		{pitch.C2, pitch.C4, pitch.C1},
		{pitch.C2, pitch.C4, pitch.C3},

		{pitch.C3, pitch.C1, pitch.C2},
		{pitch.C3, pitch.C1, pitch.C4},
		{pitch.C3, pitch.C2, pitch.C1},
		{pitch.C3, pitch.C2, pitch.C4},
		{pitch.C3, pitch.C4, pitch.C1},
		{pitch.C3, pitch.C4, pitch.C2},

		{pitch.C4, pitch.C1, pitch.C2},
		{pitch.C4, pitch.C1, pitch.C3},
		{pitch.C4, pitch.C2, pitch.C1},
		{pitch.C4, pitch.C2, pitch.C3},
		{pitch.C4, pitch.C3, pitch.C1},
		{pitch.C4, pitch.C3, pitch.C2},
	})
	testPerm(t, ps, 4, []pitch.Pitches{
		{pitch.C1, pitch.C2, pitch.C3, pitch.C4},
		{pitch.C1, pitch.C2, pitch.C4, pitch.C3},
		{pitch.C1, pitch.C3, pitch.C2, pitch.C4},
		{pitch.C1, pitch.C3, pitch.C4, pitch.C2},
		{pitch.C1, pitch.C4, pitch.C2, pitch.C3},
		{pitch.C1, pitch.C4, pitch.C3, pitch.C2},

		{pitch.C2, pitch.C1, pitch.C3, pitch.C4},
		{pitch.C2, pitch.C1, pitch.C4, pitch.C3},
		{pitch.C2, pitch.C3, pitch.C1, pitch.C4},
		{pitch.C2, pitch.C3, pitch.C4, pitch.C1},
		{pitch.C2, pitch.C4, pitch.C1, pitch.C3},
		{pitch.C2, pitch.C4, pitch.C3, pitch.C1},

		{pitch.C3, pitch.C1, pitch.C2, pitch.C4},
		{pitch.C3, pitch.C1, pitch.C4, pitch.C2},
		{pitch.C3, pitch.C2, pitch.C1, pitch.C4},
		{pitch.C3, pitch.C2, pitch.C4, pitch.C1},
		{pitch.C3, pitch.C4, pitch.C1, pitch.C2},
		{pitch.C3, pitch.C4, pitch.C2, pitch.C1},

		{pitch.C4, pitch.C1, pitch.C2, pitch.C3},
		{pitch.C4, pitch.C1, pitch.C3, pitch.C2},
		{pitch.C4, pitch.C2, pitch.C1, pitch.C3},
		{pitch.C4, pitch.C2, pitch.C3, pitch.C1},
		{pitch.C4, pitch.C3, pitch.C1, pitch.C2},
		{pitch.C4, pitch.C3, pitch.C2, pitch.C1},
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

func testPerm(t *testing.T, ps pitch.Pitches, n int, expectedPerms []pitch.Pitches) {
	name := fmt.Sprintf("perm %d from %d pitches", n, ps.Len())

	// convert pitches to strings for readability
	expected := make([][]string, len(expectedPerms))
	for i, perm := range expectedPerms {
		expected[i] = perm.Strings()
	}

	t.Run(name, func(t *testing.T) {
		actual := [][]string{}
		gatherPerms := func(perm pitch.Pitches) {
			actual = append(actual, perm.Strings())
		}

		ps.Permutation(n, gatherPerms)

		assert.ElementsMatch(t, expected, actual)
	})
}
