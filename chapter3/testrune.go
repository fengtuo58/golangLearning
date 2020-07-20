package main

import "fmt"
import "unicode/utf8"

func main() {
	var str = "hello 你好"
	fmt.Println("len str:", len(str))
	fmt.Println("RuneCountInString", utf8.RuneCountInString(str))
	fmt.Println(len([]rune(str)))
}
