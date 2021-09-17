package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	No           json.Number `json:"no"`
	Code         string      `json:"code"`
	ErrorCode    string      `json:"error_code"`
	ErrorDetail  string      `json:"error_detail"`
	ErrorMessage string      `json:"error_message"`
	Retryable    bool        `json:"retryable"`
}

//go:embed error_template.txt
var templateTxt string

const (
	templateFilePath = "error_template.txt"
	jsonFilePath     = "errors.json"
)

func main() {
	flag.Parse()
	t := template.Must(template.New("error_template").Parse(templateTxt))

	jsonData := make([]*Data, 0)

	if !strings.HasSuffix(flag.Arg(0), ".json") {
		return
	}
	b, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	for _, error := range jsonData {

		var tmp bytes.Buffer
		err = t.Execute(&tmp, error)
		if err != nil {
			log.Fatal(err)
		}

		formated, err := format.Source(tmp.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(fmt.Sprintf("errors/PG_%s.go", error.ErrorCode))
		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write(formated); err != nil {
			f.Close()
			log.Fatal(err)
		}

		f.Close()
	}

}
