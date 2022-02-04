package value

type Changes []*Change

func (cs Changes) Len() int {
	return len(cs)
}

func (cs Changes) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs Changes) Less(i, j int) bool {
	return cs[i].Offset.Cmp(cs[j].Offset) == -1
}
