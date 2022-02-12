package change

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/validation"
)

type Changes []*Change

func (changes Changes) Len() int {
	return len(changes)
}

func (changes Changes) Swap(i, j int) {
	changes[i], changes[j] = changes[j], changes[i]
}

func (changes Changes) Less(i, j int) bool {
	return changes[i].Offset.Less(changes[j].Offset)
}

func (changes Changes) Last() *Change {
	n := changes.Len()
	if n == 0 {
		return nil
	}

	return changes[n-1]
}

func (changes Changes) Validate(r ValueRange) *validation.Result {
	results := []*validation.Result{}
	errs := []error{}

	for i, change := range changes {
		if result := change.Validate(r); result != nil {
			result.Context = fmt.Sprintf("%s %d", result.Context, i)
			results = append(results, result)
		}
	}

	if len(results) == 0 && len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "changes",
		Errors:     errs,
		SubResults: results,
	}
}

// func (m Map) SortedOffsets() rat.Rats {
// 	offsets := make(rat.Rats, len(m))
// 	i := 0

// 	for offset := range m {
// 		offsets[i] = offset

// 		i++
// 	}

// 	sort.Sort(offsets)

// 	return offsets
// }
