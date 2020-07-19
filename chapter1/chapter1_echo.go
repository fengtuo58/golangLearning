package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	
	//	fmt.Println(strings.Join(os.Args[1:], "%"))
	//	fmt.Println(s)
	//fmt.Println(os.Args[0])
	for index, value := range os.Args[:] {
		fmt.Println(strconv.Itoa(index) + " " +value)
	}
}
