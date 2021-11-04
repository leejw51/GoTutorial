package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("hello world")

	x1 := new(big.Int)
	x1.SetString("100000000000000000000000000000000000000000", 10)
	y1 := new(big.Int)
	y1.SetString("3000000000000000000000000000000000000000000", 10)
	x := big.NewFloat(0).SetInt(x1)
	y := big.NewFloat(0).SetInt(y1)

	z := new(big.Float).Quo(x, y)
	z1 := new(big.Float).Quo(y, x)

	fmt.Printf("z=%f\n", z)
	fmt.Printf("z1=%f\n", z1)
	fmt.Printf("-1 z1<z cmp=%d\n", z.Cmp(z1))
	fmt.Printf("+1 z1>z cmp=%d\n", z1.Cmp(z))
}
