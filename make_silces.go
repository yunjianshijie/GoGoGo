// package main 

import "fmt"

func main() {
    a:= make([]int,5)
    fmt.Println("a",a);
    b:= make([]int,0,5) // len =0 ,cap = 5
    fmt.Println("b",b);
    c:= b[:2]
    fmt.Println("c",c);
    d:= c[2:5]
    fmt.Println("d",d)
}

