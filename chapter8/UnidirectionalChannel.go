package main

import "fmt"
var c =make(chan int)
var d =make(chan<- int)
func main(){
	fmt.Println("hello,world!")
	fmt.Println(len(c))
	d = c         // 不能写成 c = d
	<- d
	go p(d,1000)  // 虽然传给了d,但是channel按引用传递,值也传递给c了
	fmt.Println(<-c)
}

func p(ch chan<- int, i int){
	ch<-i
}
