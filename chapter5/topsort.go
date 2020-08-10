package main

import "fmt"
import "time"

func main() {
    fmt.Println("run in main goroutine")
    i := 1
    for {
        go func() {
            for {
                time.Sleep(time.Second)
            }
        }()
        if i % 10000 == 0 {
            fmt.Printf("%d goroutine started\n", i)
        }
        i++
    }
}
