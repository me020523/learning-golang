package main

import (
    "flag"
    "fmt"
    "os"
)

func demoSubcommand(){
    testSubComm := flag.NewFlagSet("test", flag.ExitOnError)
    testValue := testSubComm.String("v", "Hello,World", "some test values")
    
    testSubComm.Parse(os.Args[2:])

    fmt.Println(*testValue)
}

func main(){
  demoSubcommand()
}
