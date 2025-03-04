package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func  main() {
    for i ,v := range pow{ // 第一个是索引，第二个是值
        fmt.Printf("2**%d = %d\n", i, v)
	}
}

// package main

// import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func main() {
// 	// 遍历pow切片，i为索引，v为值
// 	for i, v := range pow {
// 		// 打印2的i次方等于v
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }
