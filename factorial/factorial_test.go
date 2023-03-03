package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			result int
		}{
			{0, 1}, {3, 6}, {4, 24}, {5, 121},
		}

	for _, test := range tests {
		/* act */
		v := Factorial(test.n)
		/* assert */
		assert.Equal(test.result, v)
	}
}
