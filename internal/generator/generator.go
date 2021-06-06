package generator

import (
	"embed"
	"fmt"
	"os"
	fp "path/filepath"
	"text/template"

	"github.com/RadhiFadlillah/dbgen/internal/sqlparser"
	"github.com/sirupsen/logrus"
)

// Generator is object to generate the code for database storage.
type Generator struct {
	// DstDir is destination directory to store the generated code.
	DstDir string
	// PackageName is the name of package for the generated code.
	PackageName string
	// TemplateFiles is files for the templates used for generating code.
	TemplateFiles embed.FS
	// DdlQueries is data of DDL queries that will be generated.
	DdlQueries []sqlparser.DdlQueryData
	// SelectQueries is data of SELECT and GET queries that will be generated.
	SelectQueries []sqlparser.SelectQueryData
	// ExecQueries is data of EXEC queries that will be generated.
	ExecQueries []sqlparser.ExecQueryData

	templates *template.Template
}

// Run execute the generator.
func (g *Generator) Run() error {
	// Prepare templates
	err := g.prepareTemplates()
	if err != nil {
		return fmt.Errorf("failed to prepare templates: %v", err)
	}

	// Make sure output directory exist
	err = os.RemoveAll(g.DstDir)
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

	err = g.generateOpenDatabase()
	if err != nil {
		return fmt.Errorf("failed to generate code for opening db: %v", err)
	}

	return nil
}

// generateAccessor generates code for database accessor.
func (g *Generator) generateAccessor() error {
	logrus.Println("generate code for accessor")

	dstPath := fp.Join(g.DstDir, "00-accessor.go")
	data := map[string]string{"package": g.PackageName}
	return g.writeCode(dstPath, "accessor.txt", data)
}

// generateOpenDatabase generates code for opening database.
func (g *Generator) generateOpenDatabase() error {
	logrus.Println("generate code for opening database")

	dstPath := fp.Join(g.DstDir, "00-database.go")
	return g.writeCode(dstPath, "database.txt", map[string]interface{}{
		"package":    g.PackageName,
		"ddlQueries": g.DdlQueries,
	})
}
