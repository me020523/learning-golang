package main

import (
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	filePath := "/Users/shuaibincheng/test.txt"

	watcher, err := fsnotify.NewWatcher()

	if(err != nil){
		log.Fatal(err)
	}
    
    waitChan := make(chan bool)
	go func(){
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event: ", event)
				if event.IsDelete(){
					waitChan <- true
				}
			case err := <-watcher.Error:
				log.Println("error: ", err)
				waitChan <- false
			}
		}
	}()

	err = watcher.Watch(filePath)
	if(err != nil){
		log.Fatal(err)
	}
	select {
	case <- waitChan:
		log.Println("exit")
	}
}