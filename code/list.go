package main

import "fmt"

type Book struct {
	title  string
	author string
}

func TestList() {
	var b []Book
	var a = Book{title: "starwars", author: "steve jobs"}
	for i := 0; i < 10; i++ {
		a.title = fmt.Sprintf("starward %d", i)
		b = append(b, a)
	}
	fmt.Printf("%v\n", b)
}
