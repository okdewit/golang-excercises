package errors

import (
	"fmt"
	"exercises/01_loops_and_functions"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	result := loops_and_functions.Sqrt(x)
	if math.IsNaN(result) {
		return 0, ErrNegativeSqrt(x)
	}
	return result, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}