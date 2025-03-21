package main

import "fmt"

type Stringer interface {
    String() string
}
type Person struct {
    Name string
    Age int
}
func  (p Person)String()string {
	return fmt.Sprintf("Name:%s, Age:%d",p.Name,p.Age)
}
// 在Go语言中，Stringer是一个接口，它定义了一个String方法，该方法返回一个字符串表示该对象。任何实现了String方法的类型都可以实现Stringer接口。

func main() {
    a := Person{"ASASA",42}
	b := Person{"BBB",11}
	
	fmt.Println(a,b)
}

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }
