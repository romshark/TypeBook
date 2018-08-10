package rend

import (
	"fmt"
	"text/template"
	"time"
)

const rendererVersion = "0.1"

type Renderer struct {
	templ *template.Template
}

// New creates a new document renderer
func New() (*Renderer, *InitStats, error) {
	// Compile HTML template
	startCompileTemplate := time.Now()
	templ, err := template.ParseFiles([]string{
		"./template/index.html",
		"./template/table-of-contents.html",
		"./template/scalar-types.html",
		"./template/enumeration-types.html",
		"./template/composite-types.html",
	}...)
	if err != nil {
		return nil, nil, fmt.Errorf("Couldn't parse template: %s", err)
	}
	compileTemplateDur := time.Since(startCompileTemplate)

	return &Renderer{
			templ: templ,
		}, &InitStats{
			CompileTemplateDur: compileTemplateDur,
		}, nil
}
