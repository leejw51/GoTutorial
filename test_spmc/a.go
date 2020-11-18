package main

import (
	"fmt"
	"time"
)

func main() {
	numberCh := make(chan int)
	
	go consumer(1, numberCh)
	go consumer(2, numberCh)
	go consumer(3, numberCh)
	producer(numberCh)
}

func producer(numberCh chan <- int) {
	for i := 0; ; i += 1 {
		fmt.Printf("produced %d\n", i)
		numberCh <- i
		<-time.After(1 * time.Second)
	}
}
		

func consumer(i int, numberCh <-chan int) {
	for {
		select {
		case task := <-numberCh:
			fmt.Printf("consumer %d receives %d\n", i, task)
		}
	}
}
