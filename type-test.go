package main

import (
	"fmt"
)

func main() {
	tp := new([7]string)
	sliceTp := tp[:]
	fmt.Printf("tp is the type of %T\n", *tp)
	fmt.Printf("sliceTp is the type of %T\n", sliceTp)
}
