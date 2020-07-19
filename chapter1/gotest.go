package main

import (
	"fmt"
	"strings"
)

func main(){
	cidr := "cerve34/44"
	ip := strings.Split(cidr, "/")
	fmt.Println("%d", len(ip))
}
