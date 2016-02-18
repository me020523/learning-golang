package main

import (
	 "fmt"
	 "encoding/binary"
)

func calc_1(){
	fmt.Printf("%d, %d\n", 9 &^ 3, 3 &^ 9)
}

func chanEqual(){
	tt := make(chan int);
	ts := tt
	tr := make(chan int);

	fmt.Println(tt == ts)
	fmt.Println(tt == tr)
}

func getLen(){
	a := int(4)
	b := uint8(4)
	c := uint32(4)
	d := int8(4)
	e := int16(4)
	f := int32(4)
	g := rune(4)
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d\n", binary.Size(&a), 
			binary.Size(b), binary.Size(c),
				 binary.Size(d), binary.Size(e), binary.Size(f),binary.Size(g) )
}
func main(){
	//calc_1()
	//chanEqual()
	getLen()
}