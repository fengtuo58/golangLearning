package main

import "fmt"
import "time"
import "runtime"
func main() {
    fmt.Println("run in main goroutine")
    runtime.GOMAXPROCS(5)
    // 读取当前的线程数
    fmt.Println(runtime.GOMAXPROCS(0))

    n := 4
	for i:=0; i<n; i++ {
        go func() {
            fmt.Println("dead loop goroutine start")
            for {}  // 死循环
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("main goroutine running")
    }
}

func topoSort(m map[string][]string) []string{
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys= append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
