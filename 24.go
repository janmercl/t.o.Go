// A Tour Of Go - 24 - Exercise: Loops and Functions

package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    z = x
    for i := 0; i < 10; i++ {
	z -= ((z*z)-x)/(2*z)	//1)
    }
    return z
}

func main() {
    a := 12.0
    fmt.Println(Sqrt(a))
    fmt.Println(math.Sqrt(a))
}
