package main

import "fmt"

type I interface {
	M()
}
type MYint int

// 定义一个函数M，参数为int类型
func (t MYint)M() {
	// do something
	// 打印参数t的值
	fmt.Println(t)
}

// 定义一个main函数
func main() {
	var i I
	 i = MYint(42)
	describe(i)
	// 调用i的M方法
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
