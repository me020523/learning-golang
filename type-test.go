package main

import (
	"fmt"
	"reflect"
)

func testSlice() {
	tp := new([7]string)
	sliceTp := tp[:]
	fmt.Printf("tp is the type of %T\n", *tp)
	fmt.Printf("sliceTp is the type of %T\n", sliceTp)
}
func testNilAndInterface(i I) {
	i.display()
}

type ts struct {
	a int
	b int
}

func testMap() {
	tm := make(map[string]int)
	a := reflect.ValueOf(tm["a"])
	fmt.Println(a)
	tm["a"] = 2
	fmt.Println(a)
}

type I interface {
	display()
}

func main() {
	testMap()
}
