package main

import "fmt"

func main() {
	s := S{"Siddharth"}
	s.display()
	var f F;
	f = 3.14
	f.display()
	fmt.Println("\nPrinting via interface object")
	var i Inter
	i = &s
	i.display()
	i = &f
	i.display()
}

type S struct {
	name string
}

type F float64

type Inter interface {
	display()
}

func (val *F) display() {
	fmt.Println(*val)
}

func (s *S) display() {
	fmt.Println(s.name)
}
