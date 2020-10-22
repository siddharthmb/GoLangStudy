package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GoLang switch practice")
	os := runtime.GOOS;

	switch os {
		case "darwin":
			fmt.Printf("Mac OSX, %v\n", os)
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("Other type: %v", os)
	}
}

