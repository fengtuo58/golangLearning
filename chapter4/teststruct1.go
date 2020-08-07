package main
import "fmt"

type ipset12 struct {
	ip string
	port int
}

type Point struct {
	X int
	Y int
}

type Circlie struct {
	Position Point
	Radios int
}

type Whelee struct {
	C Circlie
	count int
}

func main(){

	ss := make(map[ipset12]int)
	ss[ipset12{"192.168.2.112", 8080}] = 1
	wh  := new(Whelee)
	wh.count = 4
	wh.C.Position.X = 1
	wh.C.Position.Y = 2
	wh.C.Radios = 3
	d := Whelee{Circlie{Point{1,2}, 3}, 4}
	if d == *wh {
		fmt.Println("not equal")
	}
	f := Whelee{
		C: Circlie{
			Position:Point{1, 2},
			Radios:3,
		},
		count: 4,
	}
	fmt.Println(f.X)
	fmt.Printf("%#v", f)

	
}
