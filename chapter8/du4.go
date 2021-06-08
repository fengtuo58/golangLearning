package main

import(
	"fmt"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

go func() {
	os.Stdin.Read(make([]byte, 1))
	close(done)
}

for {
	select {
	case <-done:
		//drain fileSizes to allow existing goroutine to finish.
		for range fileSizes {
			
		}
		return
	case size, ok := <-fileSizes:
		//...
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		
	}
}

