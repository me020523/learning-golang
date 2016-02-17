package main

import (
    "encoding/json"
    "log"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func Marshal(){
    m := Message{"Alice", "Hello", 1294706395881547000}   
    b, err := json.Marshal(m)
    if(err != nil) {
        log.Fatal(err)
    }
    log.Println(string(b))
}
func Unmarshal(){
    jsonString := []byte(`{"name":"Bob", "body":"Hello", "time":1294706395881547000, "other":"world"}`)
    var m Message
    err := json.Unmarshal(jsonString, &m)
    if(err != nil){
        log.Fatal(err)
    }
    b, err := json.Marshal(m)
    if(err != nil){
    	log.Fatal(err)
    }
    log.Println(string(b))
}
func main() { 
    Unmarshal()
}
