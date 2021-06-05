package sqlparser

// QueryType is the type of the query.
type QueryType uint8

const (
	// DDL is query for creating table.
	DDL QueryType = iota + 1
	// SELECT is query for fetching several rows at once.
	SELECT
	// GET is query for fetching single row result.
	GET
	// EXEC is query that executed by database.
	EXEC
)

// Name returns the name of QueryType.
func (q QueryType) Name() string {
	switch q {
	case DDL:
		return "DDL"
	case SELECT:
		return "SELECT"
	case GET:
		return "GET"
	case EXEC:
		return "EXEC"
	default:
		return "UNKNOWN"
	}
}

// RawQueryData is the basic data which parsed from SQL file.
type RawQueryData struct {
	// Type is the type of query.
	Type QueryType
	// Name is the name of query.
	Name string
	// Props is list of property that defined for this query.
	Props map[string]string
	// SQL is the raw SQL query.
	SQL string
	// SourceFile is the file where this query declared.
	SourceFile string
	// SourceLine is the line number where this query declared.
	SourceLine int
}

// DdlQueryData is data for query that used to create tables.
type DdlQueryData struct {
	// TableName is the name of table that created using this DDL.
	TableName string
	// Columns is list of columns inside the created table.
	Columns []Column
	// SQL is the raw SQL query.
	SQL string
}

// Column is data of database column.
type Column struct {
	// Name is the name of the column.
	Name string
	// DbType is the database system name of the column type, e.g. VARCHAR, INT, etc.
	DbType string
	// ScanType returns a Go type suitable for scanning into using Rows.Scan.
	ScanType string
	// Nullable specify whether a column can be null or not.
	Nullable bool
}
