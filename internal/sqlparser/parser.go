package sqlparser

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Parser is the object that used to parse SQL files.
type Parser struct {
	// TmpDB is temporary database that used to test queries.
	// WARNING: make sure it's not pointed to production database.
	TmpDB *sqlx.DB
	// SrcDir is source directory that contains SQL files.
	SrcDir string

	rawQueries    []RawQueryData
	ddlQueries    []DdlQueryData
	selectQueries []SelectQueryData

	mapColumnsCount map[string]int
	mapColumnsTable map[string][]string
}

func (p *Parser) Parse() error {
	var err error

	// Parse all SQL files
	p.rawQueries, err = p.parseSqlFiles()
	if err != nil {
		err = fmt.Errorf("failed to parse sql files: %v", err)
		return err
	}

	// Process DDL queries
	p.ddlQueries, err = p.processDdlQueries()
	if err != nil {
		err = fmt.Errorf("failed to process ddl queries: %v", err)
		return err
	}

	// Store table columns to map for faster access
	p.mapColumnsTable = p.mapColumnsToTable()
	p.mapColumnsCount = p.mapTableColumnsCount()

	// Process select queries
	p.selectQueries, err = p.processSelectQueries()
	if err != nil {
		err = fmt.Errorf("failed to process select queries: %v", err)
		return err
	}

	for _, query := range p.ddlQueries {
		fmt.Println(query.TableName)
		for _, col := range query.Columns {
			fmt.Println("-", col.Name, col.DbType, col.ScanType, col.Nullable)
		}
		fmt.Println("===")
	}

	for _, query := range p.selectQueries {
		fmt.Println(query.Name)
		fmt.Println("RESULT:", query.ResultEntity)

		fmt.Println("COLUMNS:")
		for _, col := range query.Columns {
			fmt.Println("-", col.Name, col.DbType, col.ScanType, col.Nullable)
		}

		fmt.Println("PARAMETERS:")
		for _, param := range query.Params {
			fmt.Println("*", param.Name, param.Required)
		}

		fmt.Println("===")
	}

	return nil
}
