package main

import (
	"io"
	"os"
)

func main() {
	newFile, err := os.Create("copiedtext.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(newFile, f)
}
