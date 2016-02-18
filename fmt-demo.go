package main

import (
	"fmt"
)

func fmtString(){
	testStr := "Hello, Golang"
	testRunes := []rune(testStr)

	fmt.Printf("string: %s\n", testStr)
	fmt.Printf("rune: %s\n", testRunes)

	fmt.Printf("string(q): %q\n", testStr)
	fmt.Printf("rune(q): %q\n", testRunes)

	fmt.Printf("string(x): %x\n", testStr)
	fmt.Printf("rune(x): %x\n", testRunes)

	fmt.Printf("string(X): %X\n", testStr)
	fmt.Printf("rune(X): %X\n", testRunes)
}

func fmtInteger(){
	testInt := 128

	fmt.Printf("%03d\n", testInt)
	fmt.Printf("%-03#b\n", testInt)
	fmt.Printf("%-03#o\n", testInt)
	fmt.Printf("%#x\n", testInt)
	fmt.Printf("%-03#X\n", testInt)
	fmt.Printf("%-03#q\n", testInt)
	fmt.Printf("%-03#c\n", testInt)
	fmt.Printf("%-03#U\n", testInt)
}

func  fmtPointer() {
	testInt := 128
	p := &testInt
	fmt.Printf("%p\n", p)
}

func fmtStruct() {
	type MyType struct{
		Msg string
	}

	value := MyType{"Hello"}
	fmt.Printf("%d\n", value)
	fmt.Printf("%s\n", value)
	fmt.Printf("%+s\n", value)
	fmt.Printf("%+v\n", value)
}

func fmtFloating(){
	floatValue := 123.35

	fmt.Printf("%6.4G\n", floatValue)
	fmt.Printf("%.4f\n", floatValue)
}

func demoScanf(){
	var a int
	var b string

	fmt.Scanf("%d %s", &a, &b)

	fmt.Printf("a=%d, b=%s\n", a, b)
}

func main() {
	//fmtString()
	//fmtInteger()
	//fmtPointer()
	//fmtStruct()
	//fmtFloating()
	demoScanf()
}