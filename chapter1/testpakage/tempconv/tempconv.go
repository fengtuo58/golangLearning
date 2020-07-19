package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func(c Celsius) String() string{ return fmt.Sprintf("%g C", c)}
func(c Fahrenheit) String() string{ return fmt.Sprintf("%g F", c)}
