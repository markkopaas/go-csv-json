package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = ';'

	var fieldNames []string

	lineNumber := 0

	for {
		values, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		lineNumber++
		if lineNumber == 1 {
			fieldNames = values
			fmt.Println("fieldNames", fieldNames)
		} else {
			fmt.Println("data", values)
		}
	}
}
