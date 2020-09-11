package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func main2() {
	var b = []byte("helllo")
	var a = tmhash.Sum(b)
	fmt.Printf("hash = %s\n", hex.EncodeToString(a))
}

func main() {
	scores := [3]int{100, 200, 300}
	scores2 := scores
	scores[0] = 0
	fmt.Printf("%v\n", scores)
	fmt.Printf("%v\n", scores2)
}
