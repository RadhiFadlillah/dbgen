package dbgen

import (
	"fmt"
)

// mapColumnsToTable map columns to tables, which later can be used
// to match a set of columns into one of the table.
func (p *SqlParser) mapColumnsToTable(ddlQueries []DdlQueryData) map[string][]string {
	mapColumnsTable := make(map[string][]string)
	for _, ddlQuery := range ddlQueries {
		for _, col := range ddlQuery.Columns {
			colID := fmt.Sprintf("%s-%s", col.Name, col.DbType)
			mapColumnsTable[colID] = append(mapColumnsTable[colID], ddlQuery.TableName)
		}
	}

	return mapColumnsTable
}

// findSuitableTable look for table that contains all specified column.
func (p *SqlParser) findSuitableTable(columns []Column, expected string) string {
	// Find all tables that contains said columns
	tableCandidates := make(map[string]int)
	for _, col := range columns {
		colID := fmt.Sprintf("%s-%s", col.Name, col.DbType)
		for _, table := range p.mapColumnsTable[colID] {
			tableCandidates[table]++
		}
	}

	// Find tables where all columns found
	nColumn := len(columns)
	suitableTables := make(map[string]struct{})
	for tableName, count := range tableCandidates {
		if count == nColumn {
			suitableTables[tableName] = struct{}{}
		}
	}

	// If the expected is one of suitable tables, return it
	if _, exist := suitableTables[expected]; exist {
		return expected
	}

	return ""
}
