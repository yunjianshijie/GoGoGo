package main

import "golang.org/x/tour/pic"
import "math"



func Pic(dx, dy int) [][]uint8 {
	//长度为dy, [dy][dx]int
	//a  := 1
	//b  := 2
	ret := make([][]uint8,dy)
	for i:= 0 ; i < dy ;i++{
			retn := make([]uint8,dx)
		for j:=0 ;j< dx ;j++{			
			retn[j]=uint8(math.Pow(float64(i),float64(j)))
		}
		ret[i]=retn
	}
	return ret
	
}

func main() {
	pic.Show(Pic)
}
