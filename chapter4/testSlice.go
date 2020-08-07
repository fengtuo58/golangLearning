package main
import "strings"
import "fmt"
import "sort"

func remove(slice []int, i int) []int { copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	month1 := [...]string {1:"jan", 12:"dec", 4:"fff"}
	month := []string {"jan", "dec","fff"}
	//	array := [...]int {2, 3}
	fmt.Println(strings.Join(month1[:], ","))
	fmt.Println(cap(month1))
	fmt.Println(cap(month))
	s2 := "截取中文"
	f := []rune(s2)
	fmt.Println(f[:2])
	fmt.Printf("%q\n", f[0])
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	fmt.Println("KKK")
	fmt.Println(s[:0])
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	//	delete(ages, "alice")
	ages["fengtuo"]++
	fmt.Println(ages["fengtuo"])
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println("=================")
	var names []string
	for name := range ages {//迭代map 一个参数时，是map的key
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	var dd map[string]int
	fmt.Println(dd == nil)
	fmt.Println(len(dd))
	//	if _, ok := dd["dfg"]; !ok {
	//	dd["dfg"] = 9
	//}
}
