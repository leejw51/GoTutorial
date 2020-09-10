package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func main() {
	var b = []byte("helllo")
	var a = tmhash.Sum(b)
	fmt.Printf("hash = %s\n", hex.EncodeToString(a))
}
