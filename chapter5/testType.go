package main
import "fmt"
func f(...int) {}
func main() {
	fmt.Println("%T", f)
}
