package  main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) int{ // 括号是传参，外面是返回
	v.X = v.X * f
	v.Y = v.Y * f
	return 1
}
// *T 是引用的意思

func main() {
    v:= Vertex{3, 4} 
    a:= v.Scale(10)
    fmt.Println(v.Abs(),a)
}