package main

import(
	"fmt"
	"time"
)
func main() {
	var st string
	ch := make (chan int, 1)
	go func() {
		st = "123456"		
	}()
	time.Sleep(2*time.Second)
	go func() {
		fmt.Println(st)
		st = "4567"
		ch<-1
	}()
	<-ch
	fmt.Println(st)
}
