package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	gw := gzip.NewWriter(w)
	mw := io.MultiWriter(gw, os.Stdout)
	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
