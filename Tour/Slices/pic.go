package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy ; {
		ret[i] = make([]uint8, dx)
	}

	x, y := 1, 2
	for i := 0; i < dy; i++ {
		for j :=0 ; j < dx; j++ {
			ret[i][i] = uint8((x+y)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
