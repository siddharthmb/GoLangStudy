package main

import "fmt"

func fibonacci() func(int) int {
	ele := 0
	prev_val := 0
	return func(i int) int {
		if i == 0 {
			return 0
		} else if i == 1 || i == 2 {
			ele = 1
			prev_val = 1
		}else {
			ele += prev_val
			prev_val = ele - prev_val
		}
		return ele
	}
}

func main() {
	fclr := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fclr(i))
	}
}
