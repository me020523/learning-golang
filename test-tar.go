package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func writeTar(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Fatal: %s", err.Error())
		return
	}
	writer := tar.NewWriter(file)

	defer writer.Close()
	defer file.Close()

	items := [][]string{
		{"file1.txt", "this is a demo file 1"},
		{"file2.txt", "this is also a demo file 2"},
	}
	for _, i := range items {
		header := &tar.Header{
			Name: i[0],
			Mode: 0666,
			Size: int64(len(i[1])),
		}
		writer.WriteHeader(header)
		if _, err := writer.Write([]byte(i[1])); err != nil {
			fmt.Printf("Fatal: %s", err.Error())
			return
		}
	}
}

func readTar(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Fatal: %s", err.Error())
		return
	}
	reader := tar.NewReader(file)
	defer file.Close()

	for {
		header, err := reader.Next()
		if err == io.EOF {
			fmt.Println("Finishing reading.")
			break
		}
		if err != nil {
			fmt.Printf("Fatal: %s", err.Error())
			return
		}
		fmt.Printf("Name: %s\nContent:", header.Name)
		if _, err = io.Copy(os.Stdout, reader); err != nil {
			fmt.Printf("Fatal: %s", err.Error())
			return
		}
		fmt.Println()
	}
}

func main() {
	//writeTar("../test.tar")
	readTar("../test.tar")
}
