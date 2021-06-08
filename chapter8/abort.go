package main
import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commending countdown. Press return to abort")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10*time.Second):
	case <-abort:
		fmt.Println("Lauch aborted")
		return
	}
	lauch()
}


func lauch() {
	fmt.Println("Lauch ...")
}
