package main

import(
	"strings"
	"fmt"
	"bytes"
	"strconv"
)
func main() {
	fmt.Println(string(0x4eac))
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go")) // "c.d"
	fmt.Println(basename("abc")) // "abc"
	fmt.Println(comma("4387696"))
	s := "fhurh"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s2)
	fmt.Println(intsToString([]int {1, 2, 3}))
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

