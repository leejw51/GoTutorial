package engine

import "fmt"

func HackerMain() {
	fmt.Println("hacking")
	a := []int64{0, 0, 0}
	fmt.Printf("count=%d\n", len(a))
	current := len(a) - 1

	done := false
	for !done {
		a[current] = a[current] + 1
		fmt.Printf("%v\n", a)

		for j := current; j >= 0; j-- {

			if a[j] > 2 {
				a[j] = 0
				if j-1 >= 0 {
					a[j-1] = a[j-1] + 1
				} else {
					done = true
				}
			}

		}

	}
}
