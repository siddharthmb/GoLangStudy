package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = map[string]int{}
	strarr := strings.Fields(s)
	for _,str := range strarr {
		m[str]++
	}

	return m;
}

func main() {
	fmt.Println(WordCount("Demo Demo This is"))
}
