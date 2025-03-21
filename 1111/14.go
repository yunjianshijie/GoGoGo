package main

import (
	"fmt"
	//"strings"
)

func main() {
    // 声明一个空接口变量i
    var i interface{} = "hello"
	// 将i转换为字符串类型，并赋值给s
	s := i.( string ) // 类型断言 提供了访问接口值底层具体值的方式。
	// 打印s
	fmt.Println(s)

	// 将i转换为字符串类型，并赋值给s，同时判断转换是否成功，赋值给ok
	s ,ok := i.(string)
	// 打印s和ok
	fmt.Println(s,ok)

	// 将i转换为float64类型，并赋值给f，同时判断转换是否成功，赋值给ok
	f ,ok := i.(float64)
	// 打印f和ok
	fmt.Println(f,ok)

	// 将i转换为float64类型，并赋值给f，如果转换失败，会触发panic
	// f = i.(float64) // panic
	// 打印f
	fmt.Println(f)

}