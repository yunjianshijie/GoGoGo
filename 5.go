package main

import "fmt"
import "math"

type Vertex3 struct {
	X, Y float64
}

type Vertex2 struct {
	X int
	Y float64
}

func (v Vertex3) Abs() float64 {
	fmt.Printf("v is %v\n", v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 计算顶点的绝对值

func (v Vertex2) q() {
	fmt.Println(v.X, v.Y)
}

func main() {
	v := Vertex3{3, 4}
	v1:= Vertex2{3, 4}
	fmt.Println(v.Abs()) // 只接收这个类型，可以使用方法的使用
	v1.q()
}
