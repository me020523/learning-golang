package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func testWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	vals := []string{"hello", "world", "goog"}
	for _, val := range vals {
		go func(val string) {
			fmt.Println(val)
			defer wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("Done")
}
func testLocal() {
	var i uint32 = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go func() {
			defer wg.Done()
			atomic.AddUint32(&i, 1)
			fmt.Println(i)
		}()
	}
	wg.Wait()
	fmt.Println(i)
}

func main() {
	//testWaitGroup()
	testLocal()
}
