package main 

import (
	"fmt"
	engine "sprite/engine"
	fruit "sprite/engine/fruit"
	sample "sprite/engine/sample"
)

func main() {
	fmt.Println("OK")
	var a= engine.MyHelloWorld()
	fmt.Printf("%s\n", a)
	fruit.FruitHelloWorld()
	sample.MySample2()
}
