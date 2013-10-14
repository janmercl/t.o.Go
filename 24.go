package main

import (
    "fmt"
    "math"
)

var a float64

func Sqrt(x float64) float64 {
    z := 1.0
    z = x
    for i := 0; i < 10; i++ {
	z = z - ((z*z)-x)/(2*z)
    }
    return z
}

func main() {
    a = 12
    fmt.Println(Sqrt(a))
    fmt.Println(math.Sqrt(a))
}