package main

import (
	"fmt"
)

func beyondThreshold(thres int, values ...int) {
	for _, v := range values {
		if v > thres {
			fmt.Print(v)
			fmt.Print(" ")
		}
	}
}
func main() {
	beyondThreshold(10, 1, 3, 15, 4, 20)
	fmt.Println("\n")
}
