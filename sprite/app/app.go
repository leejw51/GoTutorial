package main 

import (
	"fmt"
	engine "sprite/engine"
	fruit "sprite/engine/fruit"
	sample "sprite/engine/sample"
)

func main2() {
	fmt.Println("OK")
	var a= engine.MyHelloWorld()
	fmt.Printf("%s\n", a)
	fruit.FruitHelloWorld()
	sample.MySample2()
}

func main() {
	sample.TestInterface()
}