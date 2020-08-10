
package main

import (
	"sort"
	"fmt"
)

type intSlice []string
func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func main() {
	dd := make(intSlice,0)
	dd = append(dd,"gg","guy","fd","de")
	//fmt.Println(sort.Sort(sort.Reverse(dd)))
	fmt.Println(dd)
	
	fmt.Println(sort.Reverse(dd))
	fmt.Println(dd)
}
