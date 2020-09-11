package sample

import "fmt"

type BaseCoffee struct {
	origin string
	price  int64
}

func (basecoffee *BaseCoffee) drink() {
	fmt.Printf("i drank %s coffee , and price is %d\n", basecoffee.origin, basecoffee.price)
}

type AfricaCoffee struct {
	BaseCoffee
}

func (basecoffee *AfricaCoffee) drink() {
	fmt.Printf("new!! %s coffee , and price is %d\n", basecoffee.origin, basecoffee.price)
}

func TestOverride2() {
	coffee := AfricaCoffee{
		BaseCoffee{"Brazil", 100}}
	fmt.Printf("test override2\n")
	fmt.Printf("%v\n", coffee)
	coffee.drink()
}
