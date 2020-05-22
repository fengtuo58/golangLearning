package main
import "fmt"

func main() {
	message := make (chan string)
	go func() {message <- "string"}()
	msg := <-message
	fmt.Println(msg)
}
