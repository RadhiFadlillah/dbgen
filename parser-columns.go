package dbgen

import (
	"fmt"
	"math"
	"sort"
)

// mapTableToColumnsCount map each table to its columns count.
func (p *SqlParser) mapTableToColumnsCount(ddlQueries []DdlQueryData) map[string]int {
	mapColumnsCount := make(map[string]int)
	for _, ddlQuery := range ddlQueries {
		mapColumnsCount[ddlQuery.TableName] = len(ddlQuery.Columns)
	}

	return mapColumnsCount
}

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
func (p *SqlParser) findSuitableTable(columns []Column, expectation ...string) string {
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
	var suitableTables []string
	for tableName, count := range tableCandidates {
		if count == nColumn {
			suitableTables = append(suitableTables, tableName)
		}
	}

	// If there are no suitable table, return empty
	nSuitableTables := len(suitableTables)
	if nSuitableTables == 0 {
		return ""
	}

	// If there is exactly one, return as it is
	if nSuitableTables == 1 {
		return suitableTables[0]
	}

	// If there are suitable tables, try to find the one that match with expectation
	if nSuitableTables > 0 {
		// If expected entity is one of the suitable tables, use it.
		if len(expectation) > 0 {
			expectedEntity := expectation[0]

			for _, table := range suitableTables {
				if table == expectedEntity {
					return expectedEntity
				}
			}
		}

		// If not, sort it first, then return the first one.
		sort.Slice(suitableTables, func(a, b int) bool {
			nColumnTableA := p.mapColumnsCount[suitableTables[a]]
			nColumnTableB := p.mapColumnsCount[suitableTables[b]]
			diffA := int(math.Abs(float64(nColumn - nColumnTableA)))
			diffB := int(math.Abs(float64(nColumn - nColumnTableB)))

			return diffA < diffB
		})

		return suitableTables[0]
	}

	return ""
}
