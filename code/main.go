package main

import (
	engine "testcode/engine"
	fruit "testcode/engine/fruit"
	apple "testcode/engine/fruit/apple"
)

func main() {
	engine.EngineHello()
	fruit.TestFruit()
	apple.TestApple()
}
