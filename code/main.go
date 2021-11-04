package main

import (
	"encoding/hex"
	"fmt"
	engine "testcode/engine"
	fruit "testcode/engine/fruit"
	apple "testcode/engine/fruit/apple"
)

func main2() {
	engine.EngineHello()
	fruit.TestFruit()
	apple.TestApple()
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	a, e := hex.DecodeString(encodedStr)
	if e != nil {
		b := hex.EncodeToString(a)

		fmt.Printf("%s\n", encodedStr)
		fmt.Printf("%s\n", b)
	}
	engine.JsonMain()

}

func main() {
	apple.TestApple()
}

func main4() {
	var b int64 = 2_0000_0002
	var a float64 = float64(b) / 10000_0000.0
	fmt.Printf("%d %.8f\n", b, a)
}

func main5() {
	var a int64 = 2_0000_0020
	b := a / 1_0000_0000
	c := a % 1_0000_0000
	fmt.Printf("%d.%08d\n", b, c)
}
