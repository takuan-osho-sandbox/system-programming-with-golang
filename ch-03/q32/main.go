package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test")
	if err != nil {
		log.Fatal(err)
	}
	io.CopyN(file, rand.Reader, 1024)
}
