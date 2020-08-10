package main

import "fmt"

func main() {
	test_appenddd(nil)
}

func test_appenddd(ss []string) {
	ss = append(ss, "dd")
	fmt.Printf("%v", ss)
	fmt.Println(len(ss))
	
}
