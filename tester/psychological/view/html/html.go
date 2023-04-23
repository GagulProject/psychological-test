package html

import (
	"embed"
	"io"
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

func Result(w io.Writer, p ResultParams) error {
	return result.
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
