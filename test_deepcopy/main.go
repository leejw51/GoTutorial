package main

import "fmt"

type Fruit struct {
	name  string
	price int64
	tags  []int64
}

// array, map are reference type
// string, number , struct are value type
func deepcopy(src Fruit) Fruit {
	a := make([]int64, len(src.tags))
	copy(a, src.tags)
	return Fruit{
		name:  src.name,
		price: src.price,
		tags:  a}
}

// ownership shared
func changeObj(a Fruit) {
	a.price = 200
	a.tags[0] = 200000
	fmt.Printf("changeObj a=%v\n", a)
}

func test1() {
	fmt.Printf("######### test1\n")
	a := Fruit{name: "apple", price: 100, tags: []int64{1, 2, 3}}
	fmt.Printf("a= %v\n", a)
	changeObj(a)
	fmt.Printf("a= %v\n", a)
	fmt.Printf("ownership shared, changing inside funciton affected the original\n")
}

func test2() {
	fmt.Printf("######### test2\n")
	a := Fruit{name: "apple", price: 100, tags: []int64{1, 2, 3}}
	fmt.Printf("a= %v\n", a)
	changeObj(deepcopy(a))
	fmt.Printf("a= %v\n", a)
	fmt.Printf("deepcopy, ownership not shared, changing inside funciton decoupled from the original\n")
}

func main() {
	test1()
	fmt.Printf("\n\n\n")
	test2()
}
