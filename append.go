package main 
import "fmt"
func main(){
    var s[]int
	// 可以在空白切片上追加
	s = append(s, 0)
	printS(s)
	// 在切片的末尾追加多个值
	s = append(s, 1, 2, 3)
	printS(s)
}
func printS(s []int){
    fmt.Printf("len = %d cap = %d %v\n",len(s), cap(s),s )
}