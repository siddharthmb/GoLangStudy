package main

import "fmt"
import "math"

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("\nNegative values not supported: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return  0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
