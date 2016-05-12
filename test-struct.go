package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
}
func demoFunc_1(){
	p := Person{"Tom", 15}

	b, err := json.Marshal(p)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func demoFunc_2(){
	p := Person{Age: 15}

	b, err := json.Marshal(p)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
func (p *Person) show() {
   fmt.Println(p.Name)
}
func main(){
	//demoFunc_1()
	//demoFunc_2()
        p := Person{
          Name: "Bin",
          Age: 26,
        }
        p.show()
}
