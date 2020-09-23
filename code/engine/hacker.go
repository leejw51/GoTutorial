package engine

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HackerMain() {
	fmt.Println("hacking")
	s := byte('0')
	e := byte('9')
	fmt.Print("enter sha256 hash=")
	var hash_string string = ""

	fmt.Scanln(&hash_string)
	fmt.Printf("you entered %s\n", hash_string)
	hash_bytes, _ := hex.DecodeString(hash_string)
	fmt.Printf("your hash=%x\n", hash_bytes)

	a := []byte{s, s, s, s, s, s, s, s}

	current := len(a) - 1

	done := false
	for !done {

		bs := sha256.Sum256(a)
		if 0 == bytes.Compare(hash_bytes, bs[:]) {
			fmt.Println("###############################################")
			fmt.Printf("hacked %s\n", string(a))
			fmt.Println("###############################################")
			done = true
		}
		//fmt.Printf("%v %x\n", a, bs)
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
