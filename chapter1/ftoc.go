package main
import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
	d, f := 4,6
	d, e := 4,7
	d,e,f = f,d,e
	fmt.Printf("%d %d %d", d, e, f)// simulator copy

	var p = f1()
	fmt.Printf(" %d ", *p)
	fmt.Println(f1() == f1()) // "false"
}

func f1() *int {
	v := 1
	return &v }


func fToC(f float64) float64 {
	return (f - 32)*5/9
}
