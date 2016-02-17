package main

import (
	"fmt"
)

func testA() {
	panic(1)
}

func testB() {

	defer func() {
		fmt.Println("Hello")

		a := recover()

		fmt.Println(a)
	}()

	testA()
}

func testC() {
	testB()
}

func main() {
	testC()
}
