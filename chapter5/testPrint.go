package main
import "os"
import "fmt"

func errorf(linenum int, format string, arg...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d", linenum)
	fmt.Fprintf(os.Stderr, format, arg...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	
}
