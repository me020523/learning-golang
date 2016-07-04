package main

import "fmt"

func testDefer() string {
	fmt.Println("test Defer")
	return "world"
}

func main() {
	defer testDefer()
}
