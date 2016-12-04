package main

import (
	"context"
	"log"
	"time"
)

func testCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(2) * time.Second)
		}
		log.Println("we should stop now")
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Duration(1) * time.Second)
		log.Println("hi")
	}
}

func testTimeoutContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	stop := false
	for {
		select {
		case <-ctx.Done():
			stop = true
		default:
		}
		if stop {
			log.Println("timeout")
			break
		}
		log.Println("hi")
		time.Sleep(time.Duration(1) * time.Second)
	}
	t, ok := ctx.Deadline()
	log.Printf("%v, %v\n", t, ok)
	cancel()
}

func main() {
	//testCancelContext()
	testTimeoutContext()
}
