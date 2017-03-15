package main

import (
	"fmt"
	"reflect"
)

func inspectStruct() {

	var testStruct struct {
		string
		left  int `validator:"min:1;max:10"`
		right int `validator:"min:1;max:10"`
	}
	testStruct.string = "123"
	testStruct.left = 15
	testStruct.right = 5

	t := reflect.TypeOf(testStruct)
	v := reflect.ValueOf(testStruct)
	numOfFields := t.NumField()
	fmt.Printf("Struct包含Field的个数是: %d\n", numOfFields)

	for i := 0; i < numOfFields; i++ {
		field := t.Field(i)
		value := v.Field(i)
		st := field.Tag
		fmt.Printf("%s : %s", field.Name, value.String())

		fmt.Printf(" validator: %s\n", st.Get("validator"))
	}
}
func playDeepEqual() {
	str := "123"
	ret := reflect.DeepEqual(reflect.ValueOf(str), str)
	fmt.Println(ret)
}
func customType() {
	type MyType int
	a := MyType(10)
	t := reflect.TypeOf(a)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", t.Kind())
	var b *MyType = &a
	*b = 20
	fmt.Printf("%v\n", a)
	v := reflect.ValueOf(b)
	fmt.Printf("%v\n", v)
	v.Elem().SetInt(100)
	fmt.Printf("%v", a)
}
func main() {
	//inspectStruct()
	//playDeepEqual()
	customType()
}
