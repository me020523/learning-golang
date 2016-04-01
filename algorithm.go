package main

import (
	"learning-golang/algorithm"
	"fmt"
)

func main() {
	str := "PAYPALISHIRING"
	//rowNums: = 3
	//str := "ABCDEF"
	rowNums := 9
	out := algorithm.Convert(str, rowNums)
	fmt.Println(out)	
}