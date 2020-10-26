package main

import "fmt"

type MyReader struct{}

func (e MyReader) Error() string {
	return "\nSome error occurred, study the concept again"
}

func (reader MyReader) Read(slice []byte) (int,error) {
	slice = slice[:cap(slice)]
	for i := range(slice) {
		slice[i] = 'A'
	}
    return cap(slice), nil
}

func main() {
	// this program is only meant to run on the golang playground
}
