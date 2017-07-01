package loops_and_functions

import (
	"testing"
	"math"
	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	a := assert.New(t)

	a.Equal(Sqrt(2), 1.414213562373095)
	a.Equal(Sqrt(0), 0.0)
	a.True(math.IsNaN(Sqrt(-2)))
}




