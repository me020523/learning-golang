package main

import (
	"fmt"
)

func main() {
	lambda := func(s string) { fmt.Println(s) }

	lambda("Hello,Lambda")
}
