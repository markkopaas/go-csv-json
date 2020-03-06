package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"
)

func main() {
	templatePointer := flag.String("template", "", "Output processing template. Default is json")
	flag.Parse()
	var tmpl *template.Template
	if templatePointer != nil && *templatePointer != "" {
		var err error
		tmpl, err = template.New("root").Option("missingkey=error").Parse(*templatePointer)
		if err != nil {
			panic(err)
		}
	}

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
			fieldNames = values // use the first line as field names
		} else {
			record := zipArrays(fieldNames, values)
			recordMapJson := toJson(record)
			if tmpl == nil {
				fmt.Println(recordMapJson)
			} else {
				templateData := struct {
					Record     map[string]interface{}
					RecordJson string
					LineNumber int
				}{
					record,
					recordMapJson,
					lineNumber,
				}
				err = tmpl.Execute(os.Stdout, templateData)
				if err != nil {
					panic(err)
				}
			}
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
