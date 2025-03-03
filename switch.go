package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("GO 运行的系统环境 :")
	switch os := runtime.GOOS;os{
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s",os)
	}
}
