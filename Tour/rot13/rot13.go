package main

import (
	"io"
	"os"
	"strings"
	"unicode"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (obj *rot13Reader) Read(sli []byte) (cnt int, err error) {

	cnt, err = obj.r.Read(sli)
	if err == nil {
		for i := 0; i < cnt; i++ {
			var base byte
			var ceil byte
			if unicode.IsLower(rune(sli[i])) {
				base = 97
				ceil = 122
			} else {
				base = 65
				ceil = 90
			}
			if sli[i] >= base && sli[i] < ceil {
				sli[i] = sli[i] + byte(13)
				if sli[i] > ceil {
					sli[i] = (base - 1) + (sli[i] - ceil)
				}
			}
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}
