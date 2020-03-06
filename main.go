package main

import (
	"encoding/csv"
	"encoding/json"
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
			fieldNames = values // use first line as field names
		} else {
			recordMap := zipArrays(fieldNames, values)
			recordMapJson := toJson(recordMap)
			fmt.Println(recordMapJson)
		}
	}
}

func zipArrays(keys []string, values []string) map[string]interface{} {
	recordMap := make(map[string]interface{})
	for i := 0; i < len(keys); i++ {
		recordMap[keys[i]] = values[i]
	}
	return recordMap
}

func toJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
