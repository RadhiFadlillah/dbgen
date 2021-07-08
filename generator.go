package dbgen

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	fp "path/filepath"
	"text/template"

	"github.com/sirupsen/logrus"
)

//go:embed internal/templates/*
var templateFiles embed.FS

// Generator is object to generate the code for database storage.
type Generator struct {
	// DstDir is destination directory to store the generated code.
	DstDir string
	// PackageName is the name of package for the generated code.
	PackageName string
	// DdlQueries is data of DDL queries that will be generated.
	DdlQueries []DdlQueryData
	// SelectQueries is data of SELECT and GET queries that will be generated.
	SelectQueries []SelectQueryData
	// ExecQueries is data of EXEC queries that will be generated.
	ExecQueries []ExecQueryData
	// ColumnTypeConverter is map of function to convert column's database type
	// into a suitable Go type.
	ColumnTypeConverter func(column Column) string
	// AdditionalImports is list of additional packages that should be imported
	// in the generated code, for example because of the additional types that
	// specified in ColumnTypeConverter.
	AdditionalImports []string

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

	err = g.generateInterfaces()
	if err != nil {
		return fmt.Errorf("failed to generate code for interfaces: %v", err)
	}

	err = g.generateStructs()
	if err != nil {
		return fmt.Errorf("failed to generate code for structs: %v", err)
	}

	err = g.generateSelect()
	if err != nil {
		return fmt.Errorf("failed to generate code for select queries: %v", err)
	}

	err = g.generateExec()
	if err != nil {
		return fmt.Errorf("failed to generate code for exec queries: %v", err)
	}

	// Format the generated code
	return g.formatCode()
}

// generateAccessor generates code for database accessor.
func (g *Generator) generateAccessor() error {
	logrus.Println("generate code for accessor")

	dstPath := fp.Join(g.DstDir, "accessor.go")
	data := map[string]string{"package": g.PackageName}
	return g.writeCode(dstPath, "accessor.txt", data)
}

// generateOpenDatabase generates code for opening database.
func (g *Generator) generateOpenDatabase() error {
	logrus.Println("generate code for opening database")

	dstPath := fp.Join(g.DstDir, "database.go")
	return g.writeCode(dstPath, "database.txt", map[string]interface{}{
		"package":    g.PackageName,
		"ddlQueries": g.DdlQueries,
	})
}

// generateInterfaces generates code for interfaces.
func (g *Generator) generateInterfaces() error {
	logrus.Println("generate code for interfaces")

	templateData := map[string]interface{}{
		"package":           g.PackageName,
		"ddlQueries":        g.DdlQueries,
		"selectQueries":     g.SelectQueries,
		"execQueries":       g.ExecQueries,
		"additionalImports": g.AdditionalImports,
	}

	dstPath := fp.Join(g.DstDir, "interfaces.go")
	return g.writeCode(dstPath, "interfaces.txt", templateData)
}

// generateStructs generates code for structs.
func (g *Generator) generateStructs() error {
	logrus.Println("generate code for structs")

	templateData := map[string]interface{}{
		"package":           g.PackageName,
		"ddlQueries":        g.DdlQueries,
		"selectQueries":     g.SelectQueries,
		"execQueries":       g.ExecQueries,
		"additionalImports": g.AdditionalImports,
	}

	dstPath := fp.Join(g.DstDir, "structs.go")
	return g.writeCode(dstPath, "structs.txt", templateData)
}

// generateSelect generates code for select queries.
func (g *Generator) generateSelect() error {
	logrus.Println("generate code for select queries")

	templateData := map[string]interface{}{
		"package":           g.PackageName,
		"selectQueries":     g.SelectQueries,
		"additionalImports": g.AdditionalImports,
	}

	dstPath := fp.Join(g.DstDir, "select.go")
	return g.writeCode(dstPath, "select.txt", templateData)
}

// generateExec generates code for exec queries.
func (g *Generator) generateExec() error {
	logrus.Println("generate code for exec queries")

	templateData := map[string]interface{}{
		"package":           g.PackageName,
		"execQueries":       g.ExecQueries,
		"additionalImports": g.AdditionalImports,
	}

	dstPath := fp.Join(g.DstDir, "exec.go")
	return g.writeCode(dstPath, "exec.txt", templateData)
}

// formatCode formats the generated code using goimports.
func (g *Generator) formatCode() error {
	logrus.Println("format the generated code")
	cmd := exec.Command("goimports", "-w", g.DstDir)
	return cmd.Run()
}
