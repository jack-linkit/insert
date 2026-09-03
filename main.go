package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	reader := csv.NewReader(os.Stdin)

	// read first row
	headers, err := reader.Read()
	if errors.Is(err, io.EOF) {
		log.Fatal("empty file")
		return
	}
	if err != nil {
		log.Fatalf("Error reading row: %v", err)
	}
	fmt.Printf("INSERT INTO [tablename] (%s)\n", strings.Join(headers, ", "))
	fmt.Print("VALUES ")

	rows, err := reader.ReadAll()
	if errors.Is(err, io.EOF) {
		log.Fatal("found header row but no data lines")
		return
	}
	if err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	for i, r := range rows {
		fmt.Printf("(\"%s\")", strings.Join(r, "\", \""))
		if i != len(rows)-1 {
			fmt.Print(",\n")
		}
	}

}
