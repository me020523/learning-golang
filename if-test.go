package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := make([]byte, 1024)
	os.Stdin.Read(input)
	index := bytes.IndexByte(input, '\n')
	strValue := string(input[:index])
	fmt.Println(strValue)
	value, error := strconv.Atoi(strValue)
	if error != nil {
		fmt.Println("failed to converse string to int")
		fmt.Println(error)
	}

	if fmt.Println(value); value < 60 {
		fmt.Println("D")
	} else if 60 <= value && value <= 80 {
		fmt.Println("C")
	} else if 80 < value && value < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
