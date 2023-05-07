package html

import (
	"embed"
	"strings"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	result = parse("result.html")
)

type ResultParams struct {
	Title   string
	Message string
}

func Result(p ResultParams) (*string, error) {
	var content strings.Builder
	err := result.Execute(&content, p)
	if err != nil {
		return nil, err
	}
	output := content.String()
	return &output, nil
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
