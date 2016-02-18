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
func demoFunc_3(){
	const (
		a = iota; b=3;
		c = iota
	)

	fmt.Printf("%d, %d, %d", a, b, c)
}
func main(){
	//demoFunc_1()
	//demoFunc_2();
	demoFunc_3();
}