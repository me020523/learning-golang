package main

import (
	"fmt"
	"sync"
	"time"
)

func testRepeatTimer() {
	c := make(chan int)
	go func() {
		time.AfterFunc(time.Second, func() {
			fmt.Print("expired")
			c <- 1
		})
	}()
	<-c
}
func testTicker() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan bool)
	ticker := time.NewTicker(time.Duration(100) * time.Millisecond)
	go func() {
		i := 0
		for true {
			select {
			case <-c:
				wg.Done()
				break
			case <-ticker.C:
				fmt.Println("tick")
				i++
				if i >= 10000 {
					c <- true
				}
			}
		}
	}()
	wg.Wait()
	ticker.Stop()
}
func main() {
	//	testRepeatTimer()
	testTicker()
}
