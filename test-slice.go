package main

import "fmt"

func testWriteSlice() {
	test := make([]int, 0)
	test = append(test, 1, 2, 3, 4, 5, 6)
	t1 := test[2:4]
	t1[0] = 1
	t1[1] = 2
	fmt.Printf("%v\n", test)
}
func testWriteArray() {
	var test [6]int = [6]int{1, 2, 3, 4, 5, 6}
	t1 := test[2:4]
	t1[0] = 1
	t1[1] = 2
	fmt.Printf("%v\n", test)
	t1 = append(t1, 100)
	t1[0] = 4
	fmt.Printf("%v\n", test)
	fmt.Printf("%v\n", t1)
}
func testWriteSlice2() {
	test := make([]int, 0)
	test = append(test, 1, 2, 3, 4, 5, 6)
	t1 := test[4:6]
	t1[0] = 1
	t1[1] = 2
	test = append(test, 1000)
	fmt.Printf("%v\n", test)
	t1 = append(t1, 100)
	fmt.Printf("%v\n", t1)
	fmt.Printf("%v\n", test)
}
func testSliceCap() {
	test := make([]int, 0)
	test = append(test, 1, 2, 3, 4, 5, 6)
	t1 := test[2:4]
	fmt.Printf("cap: %d\n", cap(t1))
}

func main() {
	//testWriteSlice2()
	testSliceCap()
}
