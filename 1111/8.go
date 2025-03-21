package  main
import (
	"fmt"
	"math"
)
// 接口类型 的定义为一组方法签名
type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
	f:= MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f  // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser
a =f
	// 
	//
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
