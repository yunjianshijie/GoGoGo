// # 在主文件中必须引入main的包
package main
 
import (
"fmt"
"math/rand"
"math"
)
 
// # 通过找到该main()方法进行执行程序

func main() {
        fmt.Printf("Hello,World!!!\n")
        fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
        fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
}