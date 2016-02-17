package main

import (
	"fmt"
)

func main() {
	values := []interface{}{123, "123", 'a', false, 128 + 128i}
	for _, value := range values {
		fmt.Println(value)
		switch value.(type) {
		case bool:
			fmt.Println("this is a bool")
		case int:
			fmt.Println("this is an int")
		case string:
			fmt.Println("this is a string")
		case complex128:
			fmt.Println("this is a complex")
		case rune:
			fmt.Println("this is a rune")
		default:
			fmt.Println("Unkonwn")
		}
	}
}
