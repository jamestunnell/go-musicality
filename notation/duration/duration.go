package duration

import "math/big"

type Duration struct {
	*big.Rat
}

func Zero() *Duration {
	return &Duration{Rat: big.NewRat(0, 1)}
}

func New(a, b int64) *Duration {
	return &Duration{Rat: big.NewRat(a, b)}
}

func (d *Duration) Clone() *Duration {
	return &Duration{Rat: new(big.Rat).Set(d.Rat)}
}

func (d *Duration) Zero() bool {
	return d.Rat.Cmp(Zero().Rat) == 0
}

func (d *Duration) Positive() bool {
	cmp := d.Rat.Cmp(Zero().Rat)
	return cmp == 1
}

func (d *Duration) Equal(other *Duration) bool {
	return d.Rat.Cmp(other.Rat) == 0
}

func (d *Duration) Add(other *Duration) *Duration {
	return &Duration{Rat: new(big.Rat).Add(d.Rat, other.Rat)}
}

func (d *Duration) Mul(other *Duration) *Duration {
	return &Duration{Rat: new(big.Rat).Mul(d.Rat, other.Rat)}
}
