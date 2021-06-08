package main

import(
	"fmt"
	"time"
)
func main() {
	ch := make (chan int, 0)
	var a int
	go func() {
		select {
		case ch<-3:
			fmt.Printf("Launch aborted!\n")
			return
		default:
			fmt.Println("ee")
		}
	}()
	go func() {
		a = <-ch
	}()
	time.Sleep(2*time.Second)
	fmt.Println(a)
}
