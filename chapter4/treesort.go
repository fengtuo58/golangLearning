package main

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int{
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values =appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree{
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}



func main(){
	s := "dddd"
	seen := make(map[string]struct{})
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
	} 
}

/*
struct {}
struct {}是一个无元素的结构体类型，通常在没有信息存储时使用。优点是大小为0，不需要内存来存储struct {}类型的值。
struct {} {}
struct {} {}是一个复合字面量，它构造了一个struct {}类型的值，该值也是空。
*/
