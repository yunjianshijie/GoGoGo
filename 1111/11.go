package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t // 因为t 没有赋值 

	describe(i) // (<nil>, *main.T)
	i.M()       // <nil>
	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	i.M()       // hello
	
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
