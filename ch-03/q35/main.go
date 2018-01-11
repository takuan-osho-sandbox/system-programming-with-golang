package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func CopyNx(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	written, err = io.Copy(dst, io.LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = io.EOF
	}
	return
}

func main() {
	file, err := os.Create("test")
	if err != nil {
		log.Fatal(err)
	}
	CopyNx(file, rand.Reader, 1024)
}
