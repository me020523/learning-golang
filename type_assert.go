package main

import "fmt"

func checKItType(v interface{}) {
	value, bool := v.(int)
	fmt.Println(value)
	fmt.Println(bool)
}
func main() {
	checKItType("string")
}
