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

	for {
		values, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(values)
	}
}
