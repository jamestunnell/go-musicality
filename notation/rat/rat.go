package rat

import (
	"encoding/json"
	"math/big"
)

type Rat struct {
	*big.Rat
}

var zero = big.NewRat(0, 1)

func New(a, b int64) Rat {
	return Rat{
		Rat: big.NewRat(a, b),
	}
}

func FromFloat64(x float64) Rat {
	return Rat{Rat: new(big.Rat).SetFloat64(x)}
}

func FromInt64(x int64) Rat {
	return Rat{Rat: new(big.Rat).SetInt64(x)}
}

func Zero() Rat {
	return Rat{Rat: big.NewRat(0, 1)}
}

func (r Rat) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Rat)
}

func (r *Rat) UnmarshalJSON(d []byte) error {
	var rat big.Rat

	err := json.Unmarshal(d, &rat)
	if err != nil {
		return err
	}

	r.Rat = &rat

	return nil
}

func (r Rat) Float64() float64 {
	f, _ := r.Rat.Float64()

	return f
}

func (r Rat) Clone() Rat {
	return Rat{Rat: new(big.Rat).Set(r.Rat)}
}

func (r Rat) Positive() bool {
	return r.Rat.Cmp(zero) == 1
}

func (r Rat) Negative() bool {
	return r.Rat.Cmp(zero) == -1
}

func (r Rat) Zero() bool {
	return r.Rat.Cmp(zero) == 0
}

func (r Rat) Less(other Rat) bool {
	return r.Rat.Cmp(other.Rat) == -1
}

func (r Rat) LessEqual(other Rat) bool {
	return r.Rat.Cmp(other.Rat) <= 0
}

func (r Rat) Equal(other Rat) bool {
	return r.Rat.Cmp(other.Rat) == 0
}

func (r Rat) GreaterEqual(other Rat) bool {
	return r.Rat.Cmp(other.Rat) >= 0
}

func (r Rat) Greater(other Rat) bool {
	return r.Rat.Cmp(other.Rat) == 1
}

func (r Rat) Add(other Rat) Rat {
	return Rat{Rat: new(big.Rat).Add(r.Rat, other.Rat)}
}

func (r Rat) Sub(other Rat) Rat {
	return Rat{Rat: new(big.Rat).Sub(r.Rat, other.Rat)}
}

func (r Rat) Mul(other Rat) Rat {
	return Rat{Rat: new(big.Rat).Mul(r.Rat, other.Rat)}
}

func (r Rat) MulFloat64(x float64) Rat {
	return Rat{Rat: new(big.Rat).Mul(r.Rat, new(big.Rat).SetFloat64(x))}
}

func (r Rat) MulInt64(i int64) Rat {
	return Rat{Rat: new(big.Rat).Mul(r.Rat, new(big.Rat).SetInt64(i))}
}

func (r Rat) Div(other Rat) Rat {
	return Rat{Rat: new(big.Rat).Quo(r.Rat, other.Rat)}
}

func (r Rat) Accum(other Rat) {
	r.Rat.Add(r.Rat, other.Rat)
}
