package main
import "math"
import "fmt"

var f float32 = 16777216
const e = 2.71828
const Avogadro = 6.022141129e23
const Plank = 6.62606957e-34
var z float64 
func main() {
	fmt.Println(f == f+1)

	for x:= 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	fmt.Println(z, 1/z, -1/z, z/z)
	m := math.NaN()
	fmt.Println(m == m, m <= m, m >= m)
}

