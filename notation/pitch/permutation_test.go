package pitch_test

import (
	"fmt"
	"testing"

	"github.com/jamestunnell/go-musicality/notation/pitch"
	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	testPermutation(t, 0, [][]int{
		{},
	})
	testPermutation(t, 1, [][]int{
		{0},
	})
	testPermutation(t, 2, [][]int{
		{0, 1},
		{1, 0},
	})
	testPermutation(t, 3, [][]int{
		{0, 1, 2}, {0, 2, 1},
		{1, 0, 2}, {1, 2, 0},
		{2, 0, 1}, {2, 1, 0},
	})
}

func testPermutation(t *testing.T, n int, expected [][]int) {
	t.Run(fmt.Sprintf("Permutation %d", n), func(t *testing.T) {
		actual := [][]int{}
		f := func(nums []int) {
			cpy := make([]int, len(nums))

			copy(cpy, nums)

			actual = append(actual, cpy)
		}

		pitch.Permutation(n, f)

		assert.ElementsMatch(t, expected, actual)
	})
}
