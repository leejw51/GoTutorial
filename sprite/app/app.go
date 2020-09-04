package main 

import (
	"fmt"
	sprite "example.com/sprite/x/sprite"
	fruit "example.com/sprite/x/sprite/fruit"
)

func main() {
	fmt.Println("OK")
	var a= sprite.MyHelloWorld()
	fmt.Printf("%s\n", a)
	fruit.FruitHelloWorld()
}
