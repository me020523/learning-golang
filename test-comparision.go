package main

import (
	"fmt"
)

type X interface {
	display()
}

type T struct {
	Value int
}

func (t T) display() {
	fmt.Println(t.Value)
}
func testInterfaceEqual() {
	t1 := T{
		10,
	}
	t2 := T{
		15,
	}
	var x1 X = T{
		10,
	}
	fmt.Println(t1 == t2)
	fmt.Println(t1 == x1)
	fmt.Println(t2 == x1)
}

func main() {
	testInterfaceEqual()
}
