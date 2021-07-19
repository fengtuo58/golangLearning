package main
 
import (
    "fmt"
)
 
func main() {
    fmt.Println(test())
}
 
func test() (a int) {
    a = 1
    fmt.Println("can i see ?")
	defer func() {
		if ok := recover(); ok != nil{
			a = 7
		} 
	}()
	//panic(3)
	return 2
}
