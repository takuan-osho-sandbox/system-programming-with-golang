package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name"},
		{"OCaml", "golang"},
		{"hi", "hello"},
	}

	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
