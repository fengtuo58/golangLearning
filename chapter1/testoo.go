package main

import "fmt"
func foo(value int) string {
	return ""
}

func main() {
	if str := foo(3); str == "" {
		fmt.Printf("dd")
	}
}
