package rat

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Rat struct {
	*big.Rat
}

var zero = big.NewRat(0, 1)

func New(num, denom int64) *Rat {
	return &Rat{Rat: big.NewRat(num, denom)}
}

func Zero() *Rat {
	return New(0, 1)
}

func FromFloat64(x float64) *Rat {
	return &Rat{Rat: new(big.Rat).SetFloat64(x)}
}

func FromInt64(x int64) *Rat {
	return &Rat{Rat: new(big.Rat).SetInt64(x)}
}

func FromUint64(x uint64) *Rat {
	return &Rat{Rat: new(big.Rat).SetUint64(x)}
}

func (r *Rat) String() string {
	return r.Rat.RatString()
}

func (r *Rat) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Rat) UnmarshalJSON(d []byte) error {
	var str string

	if err := json.Unmarshal(d, &str); err != nil {
		return fmt.Errorf("failed to unmarshal '%s' as string: %w", string(d), err)
	}

	rat, ok := new(big.Rat).SetString(str)
	if !ok {
		return fmt.Errorf("failed to set rat from string '%s'", str)
	}

	r.Rat = rat

	return nil
}

func (r Rat) Float64() float64 {
	f, _ := r.Rat.Float64()

	return f
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

func (r Rat) Less(other *Rat) bool {
	return r.Rat.Cmp(other.Rat) == -1
}

func (r Rat) LessEqual(other *Rat) bool {
	return r.Rat.Cmp(other.Rat) <= 0
}

func (r Rat) Equal(other *Rat) bool {
	return r.Rat.Cmp(other.Rat) == 0
}

func (r Rat) GreaterEqual(other *Rat) bool {
	return r.Rat.Cmp(other.Rat) >= 0
}

func (r Rat) Greater(other *Rat) bool {
	return r.Rat.Cmp(other.Rat) == 1
}

func (r Rat) Add(other *Rat) *Rat {
	return &Rat{Rat: new(big.Rat).Add(r.Rat, other.Rat)}
}

func (r Rat) Sub(other *Rat) *Rat {
	return &Rat{Rat: new(big.Rat).Sub(r.Rat, other.Rat)}
}

func (r Rat) Mul(other *Rat) *Rat {
	return &Rat{Rat: new(big.Rat).Mul(r.Rat, other.Rat)}
}

func (r Rat) MulFloat64(x float64) *Rat {
	return &Rat{Rat: new(big.Rat).Mul(r.Rat, new(big.Rat).SetFloat64(x))}
}

func (r Rat) MulInt64(i int64) *Rat {
	return &Rat{Rat: new(big.Rat).Mul(r.Rat, new(big.Rat).SetInt64(i))}
}

func (r Rat) MulUint64(i uint64) *Rat {
	return &Rat{Rat: new(big.Rat).Mul(r.Rat, new(big.Rat).SetUint64(i))}
}

func (r Rat) Div(other *Rat) *Rat {
	return &Rat{Rat: new(big.Rat).Quo(r.Rat, other.Rat)}
}
