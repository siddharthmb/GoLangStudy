package main

import "fmt"
import "golang.org/x/tour/tree"

func Walk (t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same (t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	m := make(map[int] int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 20; i++ {
		select {
			case v := <-ch1:
				m[v]++
			case x := <-ch2:
				m[x]++
		}
	}

	for _,val := range(m) {
		if val != 2 {
			return false
		}
	}
	
	return true
}

func main() {
//	ch := make(chan int)
//	go Walk(tree.New(1), ch)

//	for i := 0; i < 10; i++ {
//		fmt.Println(<-ch)
//	}

	if Same(tree.New(1), tree.New(10)) {
		fmt.Println("Both trees are same")
	} else {
		fmt.Println("Both trees are different")
	}
}
