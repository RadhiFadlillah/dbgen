package sqlparser

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

func (p *Parser) parseSqlFiles() ([]RawQueryData, error) {
	// Get all files in src dir
	sqlFiles, err := p.getSqlFiles()
	if err != nil {
		return nil, err
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
			rawQueries = append(rawQueries, query)
		}
	}

	return rawQueries, nil
}

func (p *Parser) parseSqlFile(path string) ([]RawQueryData, error) {
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

	return allQueries, nil
}

func (p *Parser) getSqlFiles() ([]string, error) {
	var sqlPaths []string
	err := fp.Walk(p.SrcDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && fp.Ext(path) == ".sql" {
			sqlPaths = append(sqlPaths, path)
		}
		return nil
	})

	return sqlPaths, err
}
