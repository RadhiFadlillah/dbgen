package generator

import (
	"fmt"
	"os"
	fp "path/filepath"
	"text/template"

	"github.com/sirupsen/logrus"
)

// Generator is object to generate the code for database storage.
type Generator struct {
	// DstDir is destination directory to store the generated code.
	DstDir string
	// PackageName is the name of package for the generated code.
	PackageName string
	// Templates is the templates used for generating code.
	Templates *template.Template
}

// Run execute the generator.
func (g *Generator) Run() error {
	// Make sure output directory exist
	err := os.RemoveAll(g.DstDir)
	if err != nil {
		return fmt.Errorf("failed to clean dst dir: %v", err)
	}

	os.MkdirAll(g.DstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create dst dir: %v", err)
	}

	// Generate code
	err = g.generateAccessor()
	if err != nil {
		return fmt.Errorf("failed to generate code for accessor: %v", err)
	}

	return nil
}

// generateAccessor generates code for database accessor.
func (g *Generator) generateAccessor() error {
	logrus.Println("generate accessor code")

	dstPath := fp.Join(g.DstDir, "00-accessor.go")
	data := map[string]string{"package": g.PackageName}
	return g.writeCode(dstPath, "accessor.txt", data)
}
