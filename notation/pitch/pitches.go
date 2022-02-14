package pitch

type Pitches []*Pitch

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
