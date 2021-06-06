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

	mapColumnsCount map[string]int
	mapColumnsTable map[string][]string
}

func (p *Parser) Parse() (ddlQueries []DdlQueryData, selectQueries []SelectQueryData, execQueries []ExecQueryData, err error) {
	// Parse all SQL files
	rawQueries, err := p.parseSqlFiles()
	if err != nil {
		err = fmt.Errorf("failed to parse sql files: %v", err)
		return
	}

	// Process DDL queries
	ddlQueries, err = p.processDdlQueries(rawQueries)
	if err != nil {
		err = fmt.Errorf("failed to process ddl queries: %v", err)
		return
	}

	// Store table columns to map for faster access
	p.mapColumnsTable = p.mapColumnsToTable(ddlQueries)
	p.mapColumnsCount = p.mapTableToColumnsCount(ddlQueries)

	// Process select queries
	selectQueries, err = p.processSelectQueries(rawQueries)
	if err != nil {
		err = fmt.Errorf("failed to process select queries: %v", err)
		return
	}

	// Process exec queries
	execQueries, err = p.processExecQueries(rawQueries)
	if err != nil {
		err = fmt.Errorf("failed to process exec queries: %v", err)
		return
	}

	return
}
