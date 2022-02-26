package value

type KeyMatchActionFunc func(m Map, k string, v Value) bool

// ExecuteOnKeyMatch searches the entire container recursively looking for the
// first key matching the given key, and executes the given action function.
// Recursive searching and action execution continues until the action function
// return true.
func ExecuteOnKeyMatch(v Value, keyMatch string, f KeyMatchActionFunc) bool {
	switch vv := v.(type) {
	case Map:
		for k, vvv := range vv {
			if k == keyMatch {
				if f(vv, k, vvv) {
					return true
				}
			}

			if ExecuteOnKeyMatch(vvv, keyMatch, f) {
				return true
			}
		}

	case Slice:
		for _, vvv := range vv {
			if ExecuteOnKeyMatch(vvv, keyMatch, f) {
				return true
			}
		}
	}

	return false
}
