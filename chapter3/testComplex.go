package main

import "fmt"

func main() {
	var x complex128 = complex(1, 3)
	var y complex128 = complex(2, 4)
	fmt.Println(x * y)
	fmt.Println(1i*1i)
}
