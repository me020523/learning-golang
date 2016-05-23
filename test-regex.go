package main

import (
	"fmt"
	"regexp"
)

func testSubmatchIndex(pattern, content string) {
	match := regexp.MustCompile(pattern)
	indexes := match.FindStringSubmatchIndex(content)
	fmt.Println(len(indexes))
	for _, v := range indexes {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}
func testSubmatch(pattern, content string) {
	match := regexp.MustCompile(pattern)
	submatches := match.FindAllStringSubmatch(content, -1)

	for _, row := range submatches {
		for _, m := range row {
			if len(m) > 0 {
				fmt.Print(m)
				fmt.Print(" ")
			} else {
				fmt.Print("*****")
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	//testSubmatch("\\s*(//.*)|([1-9][0-9]*)|([a-zA-Z][a-zA-Z0-9]*)", "//123459afdas29aa")
	//testSubmatch("\\s*(//.*)|([1-9][0-9]*)|([a-zA-Z][a-zA-Z0-9]*)", "123081fdasf")
	testSubmatchIndex("\\s*(//.*)|"+
		"([1-9][0-9]*)|"+
		`("(?:\\"|\\\\|[^"])*")|`+
		"("+
		"[_a-zA-Z][0-9a-zA-Z]*|"+
		"==|>=|<=|&&|\\|\\||"+
		"\\+|-|\\*|/|="+
		")", `"123\"456\"789" "fdsafds" "fdsaf\"fdsaf\\" "fsdaafds"`)
	testSubmatchIndex(`"(?:\\"|\\\\|[^"])*"`, `"123456"`)
}
