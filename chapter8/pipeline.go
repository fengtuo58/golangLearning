package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; ;x++ {
			naturals <- x
		}
	}()
	// Squrarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	//Printer(in main gorotine)
	for {
		fmt.Println(<-squares)
	}
	

}
