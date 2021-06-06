package sqlparser

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (p *Parser) processDdlQueries(rawQueries []RawQueryData) ([]DdlQueryData, error) {
	// Prepare processed queries
	var ddlQueries []DdlQueryData

	// Disable foreign keys
	logrus.Println("disabling foreign keys")
	_, err := p.TmpDB.Exec(sqlDisableForeignKey)
	if err != nil {
		err = fmt.Errorf("failed to disable foreign key: %v", err)
		return nil, err
	}

	// Create table for storing DDL query
	_, err = p.TmpDB.Exec(sqlCreateDdlCacheTable)
	if err != nil {
		err = fmt.Errorf("failed to create ddl cache table: %v", err)
		return nil, err
	}

	// Prepare statements
	stmtGetDdlCache, err := p.TmpDB.Preparex(sqlGetDdlCache)
	if err != nil {
		err = fmt.Errorf("failed to prepare stmt select: %v", err)
		return nil, err
	}

	stmtInsertDdlCache, err := p.TmpDB.Preparex(sqlInsertDdlCache)
	if err != nil {
		err = fmt.Errorf("failed to prepare stmt insert: %v", err)
		return nil, err
	}

	defer func() {
		stmtGetDdlCache.Close()
		stmtInsertDdlCache.Close()
	}()

	// Check each raw query
	for _, rawQuery := range rawQueries {
		// If it's not DDL, skip
		if rawQuery.Type != DDL {
			continue
		}

		// Fetch table name from query props
		var tableName string
		parts := rxTableName.FindStringSubmatch(rawQuery.SQL)
		if len(parts) < 2 {
			tableName = rawQuery.Props["ddl"]
		} else {
			tableName = parts[1]
		}

		// Fetch SQL for the table's DDL from cache
		var cachedDdlQuery string
		err = stmtGetDdlCache.Get(&cachedDdlQuery, tableName)
		if err != nil && err != sql.ErrNoRows {
			err = fmt.Errorf("failed to get ddl cache for table %s: %v", tableName, err)
			return nil, err
		}

		// If SQL changed, recreate table
		if rawQuery.SQL != cachedDdlQuery {
			logrus.Printf("recreating table %s", tableName)

			_, err = p.TmpDB.Exec("DROP TABLE IF EXISTS " + tableName)
			if err != nil {
				err = fmt.Errorf("failed to drop table %s: %v", tableName, err)
				return nil, err
			}

			_, err = p.TmpDB.Exec(rawQuery.SQL)
			if err != nil {
				err = fmt.Errorf("failed to create table %s: %v", tableName, err)
				return nil, err
			}

			_, err = stmtInsertDdlCache.Exec(tableName, rawQuery.SQL)
			if err != nil {
				err = fmt.Errorf("failed to save ddl cache for table %s: %v", tableName, err)
				return nil, err
			}
		}

		// Scan columns on this table
		logrus.Printf("scanning columns for %s", tableName)

		sqlSelect := "SELECT * FROM " + tableName + " LIMIT 1"
		rows, err := p.TmpDB.Query(sqlSelect)
		if err != nil && err != sql.ErrNoRows {
			err = fmt.Errorf("failed to read table %s: %v", tableName, err)
			return nil, err
		}

		columns, err := rows.ColumnTypes()
		rows.Close()
		if err != nil {
			err = fmt.Errorf("failed to scan columns for table %s: %v", tableName, err)
			return nil, err
		}

		// Create query data for this table
		query := DdlQueryData{
			TableName: tableName,
			SQL:       rawQuery.SQL,
		}

		for _, col := range columns {
			nullable, _ := col.Nullable()
			query.Columns = append(query.Columns, Column{
				Name:     col.Name(),
				DbType:   col.DatabaseTypeName(),
				ScanType: col.ScanType().String(),
				Nullable: nullable,
			})
		}

		ddlQueries = append(ddlQueries, query)
	}

	// Re-enable foreign keys
	logrus.Println("enabling foreign keys")
	_, err = p.TmpDB.Exec(sqlEnableForeignKey)
	if err != nil {
		err = fmt.Errorf("failed to enable foreign key: %v", err)
		return nil, err
	}

	return ddlQueries, nil
}
