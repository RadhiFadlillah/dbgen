package dbgen

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"

	"github.com/iancoleman/strcase"
)

var rxIn = regexp.MustCompile(`(?i)\sIN\s`)

func (g *Generator) prepareTemplates() (err error) {
	globPattern := "internal/templates/*.txt"
	funcMap := template.FuncMap{
		"camel": func(s string) string {
			return strcase.ToCamel(s)
		},
		"lowerCamel": func(s string) string {
			return strcase.ToLowerCamel(s)
		},
		"columnType": func(col Column) string {
			if g.ColumnTypeConverter == nil {
				return col.ScanType
			}
			return g.ColumnTypeConverter(col)
		},
		"useIn": func(sql string) bool {
			return rxIn.MatchString(sql)
		},
	}

	g.templates, err = template.New("dbgen").Funcs(funcMap).ParseFS(templateFiles, globPattern)
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
