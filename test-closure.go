package main

import "fmt"

func main() {
	var funcs [2]func(i int)
	for i := 0; i < 2; i++ {
		funcs[i] = func(i int) {
			fmt.Printf("Closure: %d\n", i)
		}
	}
	for i, f := range funcs {
		f(i)
	}
}
