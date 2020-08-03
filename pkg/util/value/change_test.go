package value_test

import (
	"testing"

	"github.com/jamestunnell/go-musicality/pkg/util/value"
	"github.com/stretchr/testify/assert"
)

func TestGradualChangeNonPositiveDuration(t *testing.T) {
	for _, dur := range []float64{-1, 0} {
		change, err := value.NewLinearChange(2.2, dur)

		assert.Nil(t, change)
		assert.NotNil(t, err)

		change, err = value.NewSigmoidChange(2.2, dur)

		assert.Nil(t, change)
		assert.NotNil(t, err)
	}
}
