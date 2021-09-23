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

//go:embed constructor_template.txt
var constructorTxt string

//go:embed constructor_case_template.txt
var constructorCaseTxt string

const (
	jsonFilePath = "errors.json"
)

func main() {
	flag.Parse()
	errorTemplate := template.Must(template.New("error_template").Parse(templateTxt))
	constructorTemplate := template.Must(template.New("constructor_template").Parse(constructorTxt))
	caseTemplate := template.Must(template.New("constructor_case_template").Parse(constructorCaseTxt))

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

	var allCaseBuffer bytes.Buffer

	for _, error := range jsonData {

		var tmp bytes.Buffer
		err = errorTemplate.Execute(&tmp, error)
		if err != nil {
			log.Fatal(err)
		}

		formated, err := format.Source(tmp.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		errorfile, err := os.Create(fmt.Sprintf("errors/PG_%s.go", error.ErrorCode))
		if err != nil {
			log.Fatal(err)
		}

		if _, err := errorfile.Write(formated); err != nil {
			errorfile.Close()
			log.Fatal(err)
		}

		err = caseTemplate.Execute(&allCaseBuffer, error)
		if err != nil {
			log.Fatal(err)
		}

		errorfile.Close()
	}

	var tmp bytes.Buffer
	allCase := string(allCaseBuffer.Bytes())
	err = constructorTemplate.Execute(&tmp, allCase)
	if err != nil {
		log.Fatal(err)
	}

	constructorFile, err := os.Create("errors/errors.go")
	if err != nil {
		log.Fatal(err)
	}

	formated, err := format.Source(tmp.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if _, err := constructorFile.Write(formated); err != nil {
		constructorFile.Close()
		log.Fatal(err)
	}

	constructorFile.Close()
}
