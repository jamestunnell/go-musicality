package pitch

type Pitches []*Pitch

func (ps Pitches) Clone() Pitches {
	ps2 := make(Pitches, len(ps))

	copy(ps2, ps)

	return ps2
}

func (ps Pitches) Len() int {
	return len(ps)
}

func (ps Pitches) Strings() []string {
	strs := make([]string, len(ps))
	for i, p := range ps {
		strs[i] = p.String()
	}

	return strs
}

func (ps Pitches) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps Pitches) Less(i, j int) bool {
	return ps[i].Compare(ps[j]) == -1
}

func (ps Pitches) Combination(n int, f func(comb Pitches)) {
	switch {
	case n < 0 || ps.Len() < n:
		// do nothing
	case n == 0:
		f(Pitches{})
	case n == 1:
		for _, p := range ps {
			f(Pitches{p})
		}
	default:
		for i := 1; i < ps.Len(); i++ {
			p := ps[i-1]
			ps2 := make(Pitches, ps.Len()-i)

			copy(ps2, ps[i:])

			f2 := func(comb Pitches) {
				f(append(comb, p))
			}

			ps2.Combination(n-1, f2)
		}
	}
}

func (ps Pitches) Permutation(r int, f func(perm Pitches)) {
	if r <= 0 || ps.Len() < r {
		return
	}

	ps2 := make(Pitches, r)

	ps.Combination(r, func(comb Pitches) {
		Permutation(r, func(perm []int) {
			for destIdx, srcIdx := range perm {
				ps2[destIdx] = comb[srcIdx]
			}

			f(ps2)
		})
	})
}
