package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	tmpFile, err := zipWriter.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := strings.NewReader("test")
	_, err = io.Copy(tmpFile, reader)
	if err != nil {
		log.Fatal(err)
	}
}
