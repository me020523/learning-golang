package main

import (
    "fmt"
    "os"
    "os/signal"
)

func main(){
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Kill,os.Interrupt)
    select{
	case <-sigChan:
	     fmt.Println("killed")
    }
}
