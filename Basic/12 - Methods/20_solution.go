package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	margin := 0.00000000000001
	for {
		if x < 0 {
			return 0, ErrNegativeSqrt(x)
		}
		previousZ := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(previousZ-z) < margin {
			fmt.Println(previousZ, "-", z, "=", previousZ-z)
			break
		}
	}
	fmt.Println("math.Sqrt:", math.Sqrt(x))
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
