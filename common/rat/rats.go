package rat

type Rats []Rat

func (rats Rats) Len() int {
	return len(rats)
}

func (rats Rats) Strings() []string {
	rStrings := make([]string, len(rats))

	for i, r := range rats {
		rStrings[i] = r.String()
	}

	return rStrings
}

func (rats Rats) Swap(i, j int) {
	rats[i], rats[j] = rats[j], rats[i]
}

func (rats Rats) Less(i, j int) bool {
	return rats[i].Less(rats[j])
}

func (rats Rats) Equal(other Rats) bool {
	if len(other) != len(rats) {
		return false
	}

	for i, r := range rats {
		if !r.Equal(other[i]) {
			return false
		}
	}

	return true
}

func (rats Rats) Sum() Rat {
	sum := Zero()

	for _, r := range rats {
		sum = sum.Add(r)
	}

	return sum
}

func (rats Rats) Union(other Rats) Rats {
	union := Rats{}

	union = append(union, rats...)

	for _, r := range other {
		if !union.Contains(r) {
			union = append(union, r)
		}
	}

	return union
}

func (rats Rats) Contains(other Rat) bool {
	for _, r := range rats {
		if r.Equal(other) {
			return true
		}
	}

	return false
}
