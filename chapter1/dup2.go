package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]string)	
	files := os.Args[1:]
	if len(files) == 0 {
		//		countLines(os.Stdin, counts, names)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
				continue
			}
			countLines(f, counts, names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, names[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, names map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileInfo, _ := f.Stat()
		names[input.Text()] = fileInfo.Name()
	}
}
