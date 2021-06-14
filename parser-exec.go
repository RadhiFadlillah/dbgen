package dbgen

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *SqlParser) processExecQueries(rawQueries []RawQueryData) ([]ExecQueryData, error) {
	// Prepare processed queries
	var execQueries []ExecQueryData

	// Start transaction and disable foreign keys
	tx, err := p.TmpDB.Beginx()
	if err != nil {
		err = fmt.Errorf("failed to start transaction: %v", err)
		return nil, err
	}

	defer func() {
		tx.Exec(sqlEnableForeignKey)
		tx.Rollback()
	}()

	// Disable foreign keys
	logrus.Println("disabling foreign keys")
	_, err = tx.Exec(sqlDisableForeignKey)
	if err != nil {
		err = fmt.Errorf("failed to disable foreign key: %v", err)
		return nil, err
	}

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
			logrus.Warnf("failed to execute %s: %v", queryName, err)
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
