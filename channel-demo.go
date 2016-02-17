package main
import (
    "fmt"
)

func main(){
    testChan := make(chan int)
    go func(){
        select{
	    case testChan <- 1:
		fmt.Println("heelo")
        }
    }()
    for x := range testChan {
        fmt.Println(x)	
    }
}
