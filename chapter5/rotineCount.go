package main

import "fmt"
import "time"
import "runtime"

func main() {
    fmt.Println(runtime.NumGoroutine())
    for i:=0;i<10;i++ {
        go func(){
            for {
                time.Sleep(time.Second)
            }
        }()
    }
    fmt.Println(runtime.NumGoroutine())
}
