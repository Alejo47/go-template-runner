package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"
	"text/template"
)

var jsonFlag = flag.String("json", "", "")

func init() {
	flag.Parse()
}

func main() {
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("index").Parse(string(body))
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.ReadFile(*jsonFlag)
	if err != nil {
		panic(err)
	}

	data := make(map[string]any)
	json.Unmarshal(jsonFile, &data)

	tmpl.Execute(os.Stdout, data)
}
