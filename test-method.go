package main

import "fmt"

type testStruct struct {
	a int
	b int
}

func (t *testStruct) display() {
	fmt.Printf("%d, %d\n", t.a, t.b)
}
func testMethodAsFunc() {
	t := &testStruct{
		a: 1,
		b: 2,
	}
	var f func() = t.display
	f()
}
func main() {
	testMethodAsFunc()
}
