package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func input() string {
	inByteValue := make([]byte, 1024)
	os.Stdin.Read(inByteValue)
	last := bytes.IndexByte(inByteValue, '\n')
	inByteValue = inByteValue[:last]
	return string(inByteValue)
}
func main() {
	strValue := input()
	intValue, error := strconv.Atoi(strValue)
	if error != nil {
		fmt.Println("failed to converse type string to int")
		fmt.Println(error)
		return
	}
	switch fmt.Println(intValue); intValue {
	case 10, 20, 30, 40:
		fmt.Println("A")
		fallthrough
	case 50, 60, 70, 80:
		fmt.Println("B")
	case 90, 100:
		fmt.Println("C")
	default:
		fmt.Println("Unkonwn")
	}
	return
}
