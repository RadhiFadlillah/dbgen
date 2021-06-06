package generator

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
)

func (g *Generator) prepareTemplates() (err error) {
	globPattern := "internal/templates/*.txt"
	funcMap := template.FuncMap{
		"strToCamel": func(s string) string {
			return strcase.ToCamel(s)
		},
	}

	g.templates, err = template.New("dbgen").Funcs(funcMap).ParseFS(g.TemplateFiles, globPattern)
	return
}

func (g *Generator) writeCode(dstPath, templateName string, data interface{}) error {
	// Execute template
	buffer := bytes.NewBuffer(nil)
	err := g.templates.ExecuteTemplate(buffer, templateName, data)
	if err != nil {
		return err
	}

	// Format code
	bt, err := format.Source(buffer.Bytes())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dstPath, bt, os.ModePerm)
}
