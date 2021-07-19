package main
import(
	"sync"
	"fmt"
)
func f(notify chan struct{}, wa *sync.WaitGroup) {
	wa.Add(1)
	ch := make(chan struct{})
	go f(ch, wa)
	<-notify
}

func main() {
	wait := sync.WaitGroup{}
	ch := make(chan struct{})
	wait.Add(1)
	go f(ch, &wait)
	wait.Wait()
	fmt.Print("end")
	}
