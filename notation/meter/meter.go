package meter

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jamestunnell/go-musicality/validation"
)

type Meter struct {
	Numerator, Denominator uint64
}

var meterJSONRegex = regexp.MustCompile(`"([1-9][0-9]*)/([1-9][0-9]*)"`)

func New(num, denom uint64) *Meter {
	return &Meter{
		Numerator:   num,
		Denominator: denom,
	}
}

func (m *Meter) String() string {
	return fmt.Sprintf("%d/%d", m.Numerator, m.Denominator)
}

func (m *Meter) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`"%d/%d"`, m.Numerator, m.Denominator)
	return []byte(str), nil
}

func (m *Meter) UnmarshalJSON(d []byte) error {
	results := meterJSONRegex.FindSubmatch(d)

	if results == nil {
		return fmt.Errorf("did not recognize %s as meter JSON", string(d))
	}

	numStr := string(results[1])
	denomStr := string(results[2])

	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse numerator '%s': %w", numStr, err)
	}

	denom, err := strconv.ParseUint(denomStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse denominator '%s': %w", denomStr, err)
	}

	m.Numerator = num
	m.Denominator = denom

	return nil
}

func (m *Meter) Equal(other *Meter) bool {
	return m.Numerator == other.Numerator && m.Denominator == other.Denominator
}

func (m *Meter) Validate() *validation.Result {
	errs := []error{}

	if err := validation.VerifyNonZeroUInt64("numerator", m.Numerator); err != nil {
		errs = append(errs, err)
	}

	if err := validation.VerifyNonZeroUInt64("denominator", m.Denominator); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}

	return &validation.Result{
		Context:    "meter",
		Errors:     errs,
		SubResults: []*validation.Result{},
	}
}
