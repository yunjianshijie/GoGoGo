package main
import "fmt"
func main(){
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b:= names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) //names is changed
	// 意思是类是与引用，切片是引用，数组是值
	
}