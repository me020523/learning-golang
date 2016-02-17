package main

import (
	"fmt"
)
type Showable interface {
	show()
}

type Person struct {
	name string
	age int
}

func (p Person) show() {
	//fmt.Println("name: " + p.name + ", " + "age: " + p.age)
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

type Worker struct {
	Person
	work string
}

func (w Worker) show() {
	//fmt.Println("name: " + w.name + ", age: " + w.age + ", work: " + w.work)
	//fmt.Printf("name: %s, age: %d, work: %s\n", w.name, w.age, w.work)
	p := &w
	fmt.Printf("%x, %x", w, p)

}

func newWorker(tname string, tage int, twork string) Showable {
	return Worker{
		Person: Person{
			name: tname,
			 age: tage,
		},
		work: twork,
	}
}

func main(){
	person := newWorker("Bob", 30, "Teacher")
	person.show()
}