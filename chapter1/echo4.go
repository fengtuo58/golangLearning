package main

import(
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit tailing newline")
var sep = flag.String("s", " ", "seprator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
