package main

import "fmt"

func adder() func(int)int {
	sum := 0
	return func(i int) int {
		sum += i
		fmt.Printf("In closure func sum is %d\n", sum)
		return sum
	}
}

func main() {
	val1, val2 := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(val1(i), val2(i*-10))
	}
}
