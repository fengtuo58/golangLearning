package main
import "fmt"
import "time"
func main() {
	var x, y int
	go func() {
		x = 1;
		fmt.Print("y:", y, " ")
	}()

	go func() {
		y = 1
		fmt.Print("x:", x, " ")
	}()

	time.Sleep(5*time.Second)
}
