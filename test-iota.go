package main

import "fmt"

func demoFunc_1(){
	const (
		a = iota
		b = iota
		c = iota
	)
	const d = iota
	fmt.Printf("%d, %d, %d, %d\n", a, b, c, d) //0, 1, 2, 0
}
func demoFunc_2(){
	const (
		a = iota
		b = iota
		c = 7
		d = iota
	)
	fmt.Printf("%d, %d, %d, %d\n", a,b,c,d) //0, 1, 7, 3
}
func main(){
	demoFunc_1()
	demoFunc_2();
}