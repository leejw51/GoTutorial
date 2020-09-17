package engine

import (
	"crypto/sha256"
	"fmt"
)

func HackerMain() {
	fmt.Println("hacking")
	s := byte('0')
	e := byte('9')
	a := []byte{s, s, s}

	current := len(a) - 1

	done := false
	for !done {

		bs := sha256.Sum256(a)
		fmt.Printf("%v %x\n", a, bs)
		a[current] = a[current] + 1

		for j := current; j >= 0; j-- {

			if a[j] > e {
				a[j] = s
				if j-1 >= 0 {
					a[j-1] = a[j-1] + 1
				} else {
					done = true
				}
			}

		}

	}
	fmt.Printf("count=%d  character %c~%c  byte %d~%d\n", len(a), s, e, s, e)
}
