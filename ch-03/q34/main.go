package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=test.zip")

	zipWriter := zip.NewWriter(w)
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
