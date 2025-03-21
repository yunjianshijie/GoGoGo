package main

import (
	"fmt"
	"math"
)

type MyFloat64 float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	f:=MyFloat64(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println(Abs(v))
}
