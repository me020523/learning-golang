package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func testFileConn() {
	f, err := os.Open("test.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	conn, err := net.FileConn(f) //将socket file(不是普通文件)转为Conn对象
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, _, _ := reader.ReadLine()
	fmt.Println(string(line))
}

func main() {
	testFileConn()
}
