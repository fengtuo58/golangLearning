package main
import "fmt"
func squares() func()int {
	var x int
	return func() int{
		x++
		return x*x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) //1
	fmt.Println(f())
	fmt.Println(f())
}