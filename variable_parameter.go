package main

import (
	"fmt"
)

func beyondThreshold(thres int, values ...int) {
	for _, v := range values {
		if v > thres {
			fmt.Print(v)
			fmt.Print(" ")
		}
	}
}
func calfloat(){
    if(float32(-1.0) >= float32(0)){
      fmt.Println(float32(3)/float32(2))
    }
}
func testConst(msg string) {
    const MSG = "Hello: "
//    fmt.Println(MSG + msg)
}
func testArray(count int) {
    /*val := [count]string{}
    for i := 0; i < count ; i++ {
        val[i] = "" + i
    }
    fmt.Println(val)*/
}

func main() {
	//beyondThreshold(10, 1, 3, 15, 4, 20)
//        testConst("World")
//	fmt.Println("\n")
    //testArray(10)
    calfloat()
}
