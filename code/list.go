package main

import "fmt"

type Book struct {
	title  string
	author string
}

func TestList() {
	var a = Book{title: "starwars", author: "steve jobs"}
	fmt.Printf("%v\n", a)
}
