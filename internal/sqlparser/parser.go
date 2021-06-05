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

	rawQueries []RawQueryData
}

func (p *Parser) Parse() error {
	var err error

	// Parse all SQL files
	p.rawQueries, err = p.parseSqlFiles()
	if err != nil {
		err = fmt.Errorf("failed to parse sql files: %v", err)
		return err
	}

	for _, query := range p.rawQueries {
		fmt.Println(query.Type.Name(), query.Name)
		fmt.Println(query.SQL)
		fmt.Println("===")
	}

	return nil
}
