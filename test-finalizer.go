package main

import (
	"fmt"
	"runtime"
)

type myFinalizerType struct {
	a int
}

func (a *myFinalizerType) Display() {
	fmt.Printf("a is %d\n", a.a)
}

func finalize_1() {
	obj := &myFinalizerType{
		a: 10,
	}
	runtime.SetFinalizer(obj, obj.Display)
}
func finalize_2() {
	obj := &myFinalizerType{
		a: 12,
	}
	runtime.SetFinalizer(obj, (*myFinalizerType).Display)

	obj = nil
	runtime.GC()
}
func main() {
	//finalize_1()  //出错, obj.Display不接受类型*myFinalizerType
	finalize_2()
}
