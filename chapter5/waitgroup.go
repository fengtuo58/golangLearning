package main
import "fmt"
import "sync"
func send(ch chan int, wg * sync.WaitGroup) {
	defer  wg.Done()
	i := 0
	for i <4 {
		i++
		ch <- i
	}
}

func receive(ch chan int) {
	//	value:= <- ch
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	var c = make(chan int, 4)
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	go send(c, wg)
	go send(c, wg)
	go receive(c)
	wg.Wait()
	close(c) 
}
