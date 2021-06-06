package sqlparser

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *Parser) processExecQueries(rawQueries []RawQueryData) ([]ExecQueryData, error) {
	// Prepare processed queries
	var execQueries []ExecQueryData

	// Start transaction
	tx, err := p.TmpDB.Beginx()
	if err != nil {
		err = fmt.Errorf("failed to start transaction: %v", err)
		return nil, err
	}
	defer tx.Rollback()

	// Check each raw query
	for _, rawQuery := range rawQueries {
		// If it's not EXEC, skip
		if rawQuery.Type != EXEC {
			continue
		}

		// Fetch query name
		queryName := rawQuery.Name
		logrus.Printf("processing exec query %s", queryName)

		// Find parameters in SQL query
		sqlQuery, params, queryArgs := extractQueryParams(rawQuery.SQL)

		// Try to execute query
		_, err := tx.NamedExec(sqlQuery, queryArgs)
		if err != nil {
			err = fmt.Errorf("failed to execute %s: %v", queryName, err)
			return nil, err
		}

		// Create query data for this table
		query := ExecQueryData{
			Name:   rawQuery.Name,
			SQL:    sqlQuery,
			Params: params,
		}

		execQueries = append(execQueries, query)
	}

	return execQueries, nil
}
