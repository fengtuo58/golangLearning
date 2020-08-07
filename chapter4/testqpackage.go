package main

import "p"
import "fmt"
func main() {
	var d = p.T{A: 1, B : 2}//ook
	var f = p.T{1, 2}//ok
		if d == f {
		fmt.Println("dd")
	}
}
