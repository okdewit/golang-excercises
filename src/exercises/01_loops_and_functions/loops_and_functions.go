package loops_and_functions

import "math"

func Sqrt(x float64) float64 {
	if x == 0 { return 0 }
	if x < 0 { return math.NaN() }

	z, p := x, x
	for {
		p, z = z, z - (z*z-x)/(2*z)
		if math.Abs(p - z) < 1e-15 { break }
	}

	return z
}

