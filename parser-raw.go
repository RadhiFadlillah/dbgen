package dbgen

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	fp "path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func (p *SqlParser) parseSqlFiles() ([]RawQueryData, error) {
	// Get all files in src dir
	sqlFiles, err := p.getSqlFiles()
	if err != nil {
		return nil, err
	}

	// Look for variables
	uniqueVariables := make(map[string]SqlVariable)

	for _, sqlFile := range sqlFiles {
		variables, err := p.parseVariable(sqlFile)
		if err != nil {
			return nil, err
		}

		// Make sure each query only defined once
		for _, variable := range variables {
			if _, exist := uniqueVariables[variable.Name]; exist {
				variableSource := fmt.Sprintf("%s:%d", variable.SourceFile, variable.SourceLine)
				return nil, fmt.Errorf("variable %s redefined in %s", variable.Name, variableSource)
			}

			uniqueVariables[variable.Name] = variable
		}
	}

	// Parse each SQL file
	var rawQueries []RawQueryData
	uniqueQueries := make(map[string]struct{})

	for _, sqlFile := range sqlFiles {
		queries, err := p.parseSqlFile(sqlFile)
		if err != nil {
			return nil, err
		}

		// Make sure each query only defined once
		for _, query := range queries {
			queryName := query.Type.Name() + query.Name
			if _, exist := uniqueQueries[queryName]; exist {
				querySource := fmt.Sprintf("%s:%d", query.SourceFile, query.SourceLine)
				return nil, fmt.Errorf("%s redefined in %s", query.Name, querySource)
			}
			uniqueQueries[queryName] = struct{}{}

			// Replace variables in query
			var variableErr error
			query.SQL = rxSqlVariable.ReplaceAllStringFunc(query.SQL, func(s string) string {
				parts := rxSqlVariable.FindStringSubmatch(s)
				if variable, exist := uniqueVariables[parts[1]]; exist {
					return variable.SQL
				}

				variableErr = fmt.Errorf("variable %s is not defined", parts[1])
				return s
			})

			if variableErr != nil {
				return nil, variableErr
			}

			rawQueries = append(rawQueries, query)
		}
	}

	return rawQueries, nil
}

func (p *SqlParser) parseVariable(path string) ([]SqlVariable, error) {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Prepare results
	var variable SqlVariable
	var allVariables []SqlVariable
	sqlBuffer := bytes.NewBuffer(nil)

	// Scan file line by line
	var lineNumber int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Fetch the line
		lineNumber++
		line := scanner.Text()

		// If line is SQL comment, need special care
		if strings.HasPrefix(line, "--") {
			// If SQL buffer not empty, save it first
			if variable.Name != "" && sqlBuffer.Len() > 0 {
				sql := sqlBuffer.String()
				sql = strings.TrimSpace(sql)
				sql = strings.TrimSuffix(sql, ";")

				variable.SQL = sql
				allVariables = append(allVariables, variable)
			}

			// Clean current data
			sqlBuffer.Reset()
			variable = SqlVariable{
				SourceFile: path,
				SourceLine: lineNumber,
			}

			// Find name of the variable
			for _, parts := range rxQueryProps.FindAllStringSubmatch(line, -1) {
				if parts[1] == "var" {
					variable.Name = parts[2]
					break
				}
			}

			continue
		}

		// If it's not comment and name specified, put line to buffer
		if variable.Name != "" {
			sqlBuffer.WriteString(line + "\n")
		}
	}

	// Save the last variable if exists
	if variable.Name != "" && sqlBuffer.Len() > 0 {
		sql := sqlBuffer.String()
		sql = strings.TrimSpace(sql)
		sql = strings.TrimSuffix(sql, ";")

		variable.SQL = sql
		allVariables = append(allVariables, variable)
	}

	return allVariables, nil
}

func (p *SqlParser) parseSqlFile(path string) ([]RawQueryData, error) {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Prepare results
	var query RawQueryData
	var allQueries []RawQueryData
	sqlBuffer := bytes.NewBuffer(nil)

	// Scan file line by line
	var lineNumber int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Fetch the line
		lineNumber++
		line := scanner.Text()

		// If line is SQL comment, need special care
		if strings.HasPrefix(line, "--") {
			// If query props is specified and SQL buffer not empty, save it first
			if len(query.Props) > 0 && sqlBuffer.Len() > 0 {
				sql := sqlBuffer.String()
				sql = strings.TrimSpace(sql)
				sql = strings.TrimSuffix(sql, ";")

				query.SQL = sql
				allQueries = append(allQueries, query)
			}

			// Clean current data
			sqlBuffer.Reset()
			query = RawQueryData{
				Props:      make(map[string]string),
				SourceFile: path,
				SourceLine: lineNumber,
			}

			// Find query properties
			for _, parts := range rxQueryProps.FindAllStringSubmatch(line, -1) {
				query.Props[parts[1]] = parts[2]
			}

			// Find name and type of the query.
			for key, value := range query.Props {
				value = strcase.ToCamel(value)

				switch key {
				case "ddl":
					query.Type, query.Name = DDL, "DdlCreate"+value
				case "get":
					query.Type, query.Name = GET, value
				case "select":
					query.Type, query.Name = SELECT, value
				case "exec":
					query.Type, query.Name = EXEC, value
				}

				if query.Type != 0 && query.Name != "" {
					break
				}
			}

			continue
		}

		// If it's not comment and query data not empty, put line to buffer
		if len(query.Props) > 0 {
			sqlBuffer.WriteString(line + "\n")
		}
	}

	// Save the last query if exists
	if len(query.Props) > 0 && sqlBuffer.Len() > 0 {
		sql := sqlBuffer.String()
		sql = strings.TrimSpace(sql)
		sql = strings.TrimSuffix(sql, ";")

		query.SQL = sql
		allQueries = append(allQueries, query)
	}

	return allQueries, nil
}

func (p *SqlParser) getSqlFiles() ([]string, error) {
	var sqlPaths []string
	err := fp.Walk(p.SrcDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && fp.Ext(path) == ".sql" {
			sqlPaths = append(sqlPaths, path)
		}
		return nil
	})

	return sqlPaths, err
}
