package rat

type Rats []Rat

func (rats Rats) Len() int {
	return len(rats)
}

func (rats Rats) Swap(i, j int) {
	rats[i], rats[j] = rats[j], rats[i]
}

func (rats Rats) Less(i, j int) bool {
	return rats[i].Less(rats[j])
}
