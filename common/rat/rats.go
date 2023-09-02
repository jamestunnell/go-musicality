package rat

type Rationals []*Rat

func (rats Rationals) Len() int {
	return len(rats)
}

func (rats Rationals) Strings() []string {
	rStrings := make([]string, len(rats))

	for i, r := range rats {
		rStrings[i] = r.String()
	}

	return rStrings
}

func (rats Rationals) Swap(i, j int) {
	rats[i], rats[j] = rats[j], rats[i]
}

func (rats Rationals) Less(i, j int) bool {
	return rats[i].Less(rats[j])
}

func (rats Rationals) Equal(other Rationals) bool {
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

func (rats Rationals) Sum() *Rat {
	sum := Zero()

	for _, r := range rats {
		sum = sum.Add(r)
	}

	return sum
}

func (rats Rationals) Union(other Rationals) Rationals {
	union := Rationals{}

	union = append(union, rats...)

	for _, r := range other {
		if !union.Contains(r) {
			union = append(union, r)
		}
	}

	return union
}

func (rats Rationals) Contains(other *Rat) bool {
	for _, r := range rats {
		if r.Equal(other) {
			return true
		}
	}

	return false
}
