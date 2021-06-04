package sqlparser

import (
	"bufio"
	"bytes"
	"io/fs"
	"os"
	fp "path/filepath"
	"strings"
)

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

			continue
		}

		// If it's not comment and query data not empty, put line to buffer
		if len(query.Props) > 0 {
			sqlBuffer.WriteString(line + "\n")
		}
	}

	return allQueries, nil
}
