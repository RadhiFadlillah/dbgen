package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/RadhiFadlillah/dbgen/internal/sqlparser"
	"github.com/iancoleman/strcase"
)

func (g *Generator) prepareTemplates() (err error) {
	globPattern := "internal/templates/*.txt"
	funcMap := template.FuncMap{
		"camel": func(s string) string {
			return strcase.ToCamel(s)
		},
		"lowerCamel": func(s string) string {
			return strcase.ToLowerCamel(s)
		},
		"columnType": func(col sqlparser.Column) string {
			if g.ColumnTypeConverter == nil {
				return col.ScanType
			}
			return g.ColumnTypeConverter(col)
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

	return ioutil.WriteFile(dstPath, buffer.Bytes(), os.ModePerm)
}
