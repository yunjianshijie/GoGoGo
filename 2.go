package main

import (
	"fmt"
	"math"
	// "math/rand"
	//"math"
	"math/cmplx"
)

var (
	T bool       = true
	F bool       = false
	z complex128 = cmplx.Sqrt(5 + 12i)
)

const Pi = 3.14 // := 也不能用于声明常量
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	//再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func main() {
	//ar i int
	//var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f) //
	// // var b bool
	// // var s string
	// // x := 34 不行因为第一次是声明 第二次是赋值

	// j:= Pi
	// fmt.Println(x, y, z,j)
	// //fmt.Printf("%v %v %v %q\n", i, f, b, s) // 副0值
	// fmt.Printf("类型：%T\n z的值为:%v\n", z, z)
	// fmt.Printf("true && false = ", T && F)

	// 调用needInt函数，传入Small参数，输出结果
	fmt.Println(needInt(Small))
	// 调用needFloat函数，传入Small参数，输出结果
	fmt.Println(needFloat(Small))
	// 调用needFloat函数，传入Big参数，输出结果
	fmt.Println(needFloat(Big))

	// 声明变量sum，初始值为0
	sum := 0
	// for循环，从0开始，每次递增1，直到sum小于10
	for i := 0; i < 10; i++ {
		// sum加上i
		sum += i
	}
	// for循环，从sum小于1000开始，每次sum乘以2
	for sum < 1000 {
		// sum加上sum
		sum += sum
	}
	// for sum < 1000 {
	// 	sum += sum
	// }
	// 这就是while

	fmt.Println(sum)
		fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow( 2, 3, 10),pow(3,3,20))
}

func pow(x,n,lim float64 )float64 {
	if v:= math.Pow(x,n); v < lim {
		return v
	}else{
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}

fu

