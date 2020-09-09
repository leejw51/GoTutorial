package sample

import "fmt"

type Fruit struct {
	price int
	kind  int
	name  string
}

func change(output *Fruit) {
	output.name = "pear"
}
func TestPointer() {
	//	sample.TestInterface()
	var a = Fruit{price: 0, kind: 1, name: "apple"}
	change(&a)
	fmt.Println(a)
}
