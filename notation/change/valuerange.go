package change

import "fmt"

type ValueRange interface {
	Includes(float64) bool
	String() string
}

type MinMaxInclRange struct {
	Min, Max float64
}

func (r *MinMaxInclRange) Includes(v float64) bool {
	return v >= r.Min && v <= r.Max
}

func (r *MinMaxInclRange) String() string {
	return fmt.Sprintf("[%v,%v]", r.Min, r.Max)
}

type MinExclRange struct {
	Min float64
}

func (r *MinExclRange) Includes(v float64) bool {
	return v > r.Min
}

func (r *MinExclRange) String() string {
	return fmt.Sprintf("(%v,inf)", r.Min)
}
