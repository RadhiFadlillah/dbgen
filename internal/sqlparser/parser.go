package sqlparser

import (
	"github.com/jmoiron/sqlx"
)

// Parser is the object that used to parse SQL files.
type Parser struct {
	// TmpDB is temporary database that used to test queries.
	// WARNING: make sure it's not pointed to production database.
	TmpDB *sqlx.DB
	// SrcDir is source directory that contains SQL files.
	SrcDir string

	rawQueries []RawQueryData
}

func (p *Parser) Parse() error {
	// Get all files in src dir
	sqlFiles, err := p.getSqlFiles()
	if err != nil {
		return err
	}

	// Parse each SQL file
	var rawQueries []RawQueryData
	for _, sqlFile := range sqlFiles {
		queries, err := p.parseSqlFile(sqlFile)
		if err != nil {
			return err
		}

		rawQueries = append(rawQueries, queries...)
	}

	return nil
}
