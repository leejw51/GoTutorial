package main

import (
	"encoding/hex"
	"fmt"
	engine "testcode/engine"
	fruit "testcode/engine/fruit"
	apple "testcode/engine/fruit/apple"
)

func main() {
	engine.EngineHello()
	fruit.TestFruit()
	apple.TestApple()
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)
}
