package main

import "fmt"

func add(x int, y int)int {
	return x + y
}
func sub(x, y int) (z int) {
	z=x-y;return
}

func first(x int, _ int) int{
	return x
}

func zero(int, int) int {
	return 0
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main() {
	fmt.Printf("%*s%T\n", 7, " ",add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
	if 1 > 2 || (2 != 3 && 3 != 4) {
		fmt.Println("ok")
	}
	f := squares()
	fmt.Println(f()) //"1"
	fmt.Println(f()) //"4"
	fmt.Println(f()) //"9"
}
