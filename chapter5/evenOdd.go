package main

import "fmt"

func main() {
	ch0 := make(chan int)
	ch1 := make(chan int)
	cxtDone := make(chan int)	
	go func() {
		for i := 1; i < 10; i+=2 {
			<- ch0
			fmt.Println(i)
			ch1 <- 0
			if i == 9 {
				break
			}

		}
	}()

	go func() {
		for i := 2; i <= 10; i+=2 {
			<-ch1
			fmt.Println(i)
			if i == 10 {
				cxtDone <- 0
				break
			}
			ch0 <-0
		}
	}()

	ch0 <- 0
	select {
	case <- cxtDone:
		fmt.Println("print ok")
	}
}
