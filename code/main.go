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
}

func main() {
	engine.JsonMain()
}
