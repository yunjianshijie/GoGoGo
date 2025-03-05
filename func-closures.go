package main

import "fmt"

// 定义一个函数adder，返回一个函数，该函数接受一个整数参数，返回一个整数
func adder()func (int)int{
	// 定义一个变量sum，初始值为0
	sum := 0
	// 返回一个匿名函数，该函数接受一个整数参数x，将sum加上x，并返回sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			adder()(i), // 这是为什么？ 因adder返回函数但是adder运行了一边，但是pos=adder(),每次运行sum都变了
			pos(i),
			neg(-2*i),
		)
	}
}