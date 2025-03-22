// //Readers

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
//     r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)

// 	for {
// 	    n, err := r.Read(b)
// 		fmt.Printf("n= %v err = %v b = %v\n",n,err,b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])

// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

package main

import (
	"fmt"
	//"io"
)

// Reader 结构体
type Reader struct{}

// Read 方法实现 io.Reader 接口
func (r *Reader) Read(p []byte) (n int, err error) {
	// 产生一个 ASCII 字符 'A'
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil // 返回读入的字节数
}

func main() {
	reader := &Reader{}        // 创建一个 Reader 实例
	buffer := make([]byte, 10) // 创建一个缓冲区
	for {

		n, err := reader.Read(buffer) // 从 Reader 读取数据
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Print(string(buffer[:n])) // 打印读取的内容打打印前n个字节
	}
}
