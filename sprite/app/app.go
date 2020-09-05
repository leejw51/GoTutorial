package main 

import (
	"fmt"
	engine "sprite/engine"
	fruit "sprite/engine/fruit"
)

func main() {
	fmt.Println("OK")
	var a= engine.MyHelloWorld()
	fmt.Printf("%s\n", a)
	fruit.FruitHelloWorld()
}
