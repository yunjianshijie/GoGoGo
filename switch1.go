package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Println("周六是哪天？")
// 	today := time.Now().Weekday()
// 	fmt.Println(today+1)
// 	switch time.Saturday { //
// 	case today + 0:
// 		fmt.Println("今天")
// 	case today + 1:
// 		fmt.Println("明天")
// 	case today + 2:
// 		fmt.Println("昨天")
// 	default:
// 		fmt.Println("太远了")
// 	}
// }


func main(){
	t:= time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("早上好")
	case t.Hour()<17:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}
}