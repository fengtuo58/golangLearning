package main

import "testpackage/tempconv"
import "fmt"

func main () {
	fmt.Printf("%d", tempconv.CToF(tempconv.BoilingC))
}

func init () {
	fmt.Printf("inti!!\n") //construct
	}
