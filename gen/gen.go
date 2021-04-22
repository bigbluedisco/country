package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type Info struct {
	Name              string `json:"name"`
	ISO3166_1_numeric int    `json:"iso_3166_1_numeric"`
	ISO3166_2         string `json:"iso_3166_2"`
	ISO3166_3         string `json:"iso_3166_3"`
	DefaultCurrency   string `json:"default_currency"`
	DefaultLanguage   string `json:"default_language"`
}

var countryTemplate = template.Must(template.New("").Parse(`
		"{{ .ISO3166_2 }}": {
			Name:              "{{ .Name }}",
			ISO3166_1_numeric: {{ .ISO3166_1_numeric }},
			ISO3166_2:         "{{ .ISO3166_2 }}",
			ISO3166_3:         "{{ .ISO3166_3 }}",
			DefaultCurrency:   "{{ .DefaultCurrency }}",
			DefaultLanguage:   "{{ .DefaultLanguage }}",
		},`))

func main() {

	bt, err := ioutil.ReadFile("./countries.json")
	if err != nil {
		log.Fatal(err)
	}

	var countries []Info
	if err := json.Unmarshal(bt, &countries); err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	for _, c := range countries {
		if err := countryTemplate.Execute(&b, c); err != nil {
			log.Fatal(c, err)
		}
	}

	if err := ioutil.WriteFile("out.go", []byte(b.String()), 0644); err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
