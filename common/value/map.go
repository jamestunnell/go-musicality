package value

type Map map[string]Value

type MapModFunc func(m Map)

// FindValue searches the entire container recursively looking for the
// first key matching the given key, and retrieves the value.
func (rootMap Map) FindValue(matchKey string) (Value, bool) {
	var val Value
	f := func(parentMap Map, k string, v Value) bool {
		val = v

		return true
	}

	return val, ExecuteOnKeyMatch(rootMap, matchKey, f)
}

// ChangeValue searches the entire container recursively looking for the
// first key matching the given key, and changes the associated value in the parent map.
func (rootMap Map) ChangeValue(matchKey string, newValue Value) bool {
	f := func(parentMap Map, k string, v Value) bool {
		parentMap[k] = newValue

		return true
	}

	return ExecuteOnKeyMatch(rootMap, matchKey, f)
}

// RemoveValue searches the entire container recursively looking for the
// first key matching the given key, and removes the key from the parent map.
func (rootMap Map) RemoveValue(matchKey string) bool {
	f := func(parentMap Map, k string, v Value) bool {
		delete(parentMap, matchKey)

		return true
	}

	return ExecuteOnKeyMatch(rootMap, matchKey, f)
}

func (rootMap Map) ChangeSliceValue(sliceKey string, index int, newValue Value) bool {
	f := func(parentMap Map, k string, v Value) bool {
		s, ok := v.(Slice)
		if !ok {
			return false
		}

		s[index] = newValue
		// parentMap[k] = newValue

		return true
	}

	return ExecuteOnKeyMatch(rootMap, sliceKey, f)
}

func (rootMap Map) GetSliceValue(sliceKey string, index int) (Value, bool) {
	var val Value

	f := func(parentMap Map, k string, v Value) bool {
		s, ok := v.(Slice)
		if !ok {
			return false
		}

		val = s[index]

		return true
	}

	if ExecuteOnKeyMatch(rootMap, sliceKey, f) {
		return val, true
	}

	return nil, false
}
