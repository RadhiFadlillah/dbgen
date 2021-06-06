package generator

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
)

func (g *Generator) writeCode(dstPath, templateName string, data interface{}) error {
	// Execute template
	buffer := bytes.NewBuffer(nil)
	err := g.Templates.ExecuteTemplate(buffer, templateName, data)
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
