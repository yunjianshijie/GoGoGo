package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
} // 这里函数func  是一个参数


func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//
	fmt.Println(hypot) // 0x492160 hypot像函数指针 
	fmt.Println(hypot(5, 12))
 
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
