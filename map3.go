// package main

// import "fmt"

// type Vertex struct {
//     Lat, Long float64
// }

// var m map[string]Vertex{
//     "Bell Labs": {    40.68433, -74.39967, },
//     "Google": {
//         37.42202, -122.08408,
//     },
// }

// func main() {
//     fmt.Println(m)
// }

package main
import "fmt"
func main(){
	m:= make(map[string]int)
	m["one"]=1
	m["two"]=2
	fmt.Println("m:",m["one"])
	delete(m,"one")
	fmt.Println("m:",m["one"])
	v,ok:=m["one"]
	fmt.Println("v:",v,"ok:",ok)
}