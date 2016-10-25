package main

import (
  "path/filepath"
  "fmt"
)

func main(){
    ok, err := filepath.Match("/usr/bin*/123", "/usr/bin123/123")
    fmt.Println(ok)
    fmt.Println(err.Error())
}
