package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type Info struct {
	Name              string   `json:"name"`
	ISO3166_1_numeric int      `json:"iso_3166_1_numeric"`
	ISO3166_2         string   `json:"iso_3166_2"`
	ISO3166_3         string   `json:"iso_3166_3"`
	DefaultCurrency   string   `json:"default_currency"`
	DefaultLanguage   string   `json:"default_language"`
	Languages         []string `json:"languages"`
}

var countryTemplate = template.Must(template.New("").Parse(`
		"{{ .ISO3166_2 }}": {
			Name:              "{{ .Name }}",
			ISO3166_1_numeric: {{ .ISO3166_1_numeric }},
			ISO3166_2:         "{{ .ISO3166_2 }}",
			ISO3166_3:         "{{ .ISO3166_3 }}",
			DefaultCurrency:   "{{ .DefaultCurrency }}",
			DefaultLanguage:   "{{ .DefaultLanguage }}",
			Languages: []string{
				{{- range $val := .Languages }}
				"{{- $val }}",
			{{- end }}
			},
		},`))

func main() {
	lt, err := ioutil.ReadFile("./lang.json")
	if err != nil {
		log.Fatal(err)
	}

	var langs map[string][]string
	if err := json.Unmarshal(lt, &langs); err != nil {
		log.Fatal(err)
	}

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

		// Fill up languages from the other source if the primary source did not privide any.
		if len(c.Languages) == 0 {
			if lgs, ok := langs[c.ISO3166_2]; ok {
				c.DefaultLanguage = lgs[0]
				c.Languages = lgs
			}
		}

		if err := countryTemplate.Execute(&b, c); err != nil {
			log.Fatal(c, err)
		}
	}

	if err := ioutil.WriteFile("out.go", []byte(b.String()), 0644); err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}
