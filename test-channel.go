package main

import (
	"fmt"
)

func testNilSend() {
	var testChan = make(chan interface{}, 1)

	go func() {
		fmt.Println("push a test value")
		testChan <- nil
	}()

	val := <-testChan

	fmt.Println(val)
}

func main() {
	testNilSend()
}
