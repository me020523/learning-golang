package main

import (
	"fmt"
)

func testChanRange() {
	testChan := make(chan int)
	go func() {
		select {
		case testChan <- 1:
			fmt.Println("heelo")
		}
	}()
	for x := range testChan {
		fmt.Println(x)
	}
}

func main() {
	testChanRange()
}
