package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	fmt.Sprint(float64(e))
	return fmt.Sprintf("Cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		if math.Abs((z*z-x)/(2*z)) < 0.0000001 {
			return z, nil
		}
		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
	}
	return z, nil
}

func main() {
	//fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	//fmt.Println(ErrNegativeSqrt(-2))
}
