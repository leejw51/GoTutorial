package main

import (
	"fmt"
	engine "sprite/engine"
	fruit "sprite/engine/fruit"
	sample "sprite/engine/sample"
)

func main2() {
	fmt.Println("OK")
	var a = engine.MyHelloWorld()
	fmt.Printf("%s\n", a)
	fruit.FruitHelloWorld()
	sample.MySample2()
}

type Fruit struct {
	price int
	kind  int
	name  string
}

func change(output *Fruit) {
	output.name = "pear"
}
func main() {
	//	sample.TestInterface()
	var a = Fruit{price: 0, kind: 1, name: "apple"}
	change(&a)
	fmt.Println(a)
}
