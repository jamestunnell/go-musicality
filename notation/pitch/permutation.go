package pitch

// Generates permutations of numbers from 0 to n-1, passing them to the given func.
// The permutations reuse the same array for storage, so they must be used or copied
// when received by the function.
func Permutation(n int, f func([]int)) {
	switch {
	case n < 0:
		// do nothing
	case n == 0:
		f([]int{})
	default:
		var finished bool

		result := make([]int, n)

		for i := range result {
			result[i] = i
		}

		f(result)

		for {
			finished = true

			for i := n - 1; i > 0; i-- {

				if result[i] > result[i-1] {
					finished = false

					minGreaterIndex := i
					for j := i + 1; j < n; j++ {
						if result[j] < result[minGreaterIndex] && result[j] > result[i-1] {
							minGreaterIndex = j
						}

					}

					result[i-1], result[minGreaterIndex] = result[minGreaterIndex], result[i-1]

					//sort from i to n-1
					for j := i; j < n; j++ {
						for k := j + 1; k < n; k++ {
							if result[j] > result[k] {
								result[j], result[k] = result[k], result[j]
							}

						}
					}
					break
				}
			}

			if finished {
				break
			}

			f(result)
		}
	}
}
