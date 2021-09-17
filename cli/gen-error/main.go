package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
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
		f, err := os.Create(fmt.Sprintf("errors/PG_%s.go", error.ErrorCode))
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(f, error)
		if err != nil {
			f.Close()
			log.Fatal(err)
		}

		f.Close()
	}

}