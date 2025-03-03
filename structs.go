package main
import "fmt"

type Vertex struct {
	X ,Y int	
}


func main(){

	// fmt.Println(Vertex{1,2})
	 v := Vertex{1,2}
	// v.X = 4
	v3 := Vertex{X:1} //  Y is 0 
	p := &v
	p.X = 1e9
	fmt.Println(v,v3)
	fmt.Println(p)
}