package sqlparser

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *Parser) processSelectQueries() ([]SelectQueryData, error) {
	// Prepare processed queries
	var selectQueries []SelectQueryData

	// Check each raw query
	for _, rawQuery := range p.rawQueries {
		// If it's not SELECT or GET, skip
		if rawQuery.Type != SELECT && rawQuery.Type != GET {
			continue
		}

		// Fetch query name
		queryName := rawQuery.Name
		logrus.Printf("processing select query %s", queryName)

		// Find parameters in SQL query
		params, queryArgs := extractQueryParams(rawQuery.SQL)

		// Execute select query
		rows, err := p.TmpDB.NamedQuery(rawQuery.SQL, queryArgs)
		if err != nil && err != sql.ErrNoRows {
			err = fmt.Errorf("failed to run %s: %v", queryName, err)
			return nil, err
		}

		// Get the columns
		columns, err := rows.ColumnTypes()
		rows.Close()
		if err != nil {
			err = fmt.Errorf("failed to get columns for query %s: %v", queryName, err)
			return nil, err
		}

		// Create query data for this table
		query := SelectQueryData{
			Name:         rawQuery.Name,
			SQL:          rawQuery.SQL,
			Params:       params,
			SingleResult: rawQuery.Type == GET,
		}

		// Save columns for this query
		for _, col := range columns {
			query.Columns = append(query.Columns, Column{
				Name:     col.Name(),
				DbType:   col.DatabaseTypeName(),
				ScanType: col.ScanType().String(),
			})
		}

		// Find suitable entity to store result.
		expectedResultEntity := rawQuery.Props["result"]
		query.ResultEntity = p.findSuitableTable(query.Columns, expectedResultEntity)

		selectQueries = append(selectQueries, query)
	}

	return selectQueries, nil
}
