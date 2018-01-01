package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "%d is int and %f is float. %s is string", 1, 1.0, "一")
}
