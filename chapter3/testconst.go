package main

import "fmt"

const IPv4Len = 4

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednessday
	Thursday
	Friday
	Saturday
)

const (
	deadbeef = 0xdeadbeef
	a = uint32(deadbeef)
	b = float32(deadbeef)
	c = float64(deadbeef)
	//	d = int32(deadbeef)
	//	e = float64(1e300)
	//	f = uint(-1)
)

type Flags uint 
const (
	top Flags = 1 <<iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

type Power struct{
        age int
        high int
        name string
}

func parseIpv4() {
	var p [IPv4Len]byte
	p[0] = 8	
}

func main() {
	//	fmt.Printf("%b \n", FlagBroadcast)
	var i Power = Power{age: 10, high: 178, name: "NewMan"}
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')

	

        fmt.Printf("type:%T\n", i)
        fmt.Printf("value:%v\n", i)
        fmt.Printf("value+:%+v\n", i)
        fmt.Printf("value#:%#v\n", i)


}


