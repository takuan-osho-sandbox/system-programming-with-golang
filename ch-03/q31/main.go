package main

import (
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.OpenFile("new.txt", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}
}
