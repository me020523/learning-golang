package main

import (
	"fmt"
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

type I interface {
	display()
}

func main() {
	var i I = nil
	testNilAndInterface(i)
}
