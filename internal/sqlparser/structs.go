package sqlparser

// RawQueryData is the basic data which parsed from SQL file.
type RawQueryData struct {
	// Props is list of property that defined for this query.
	Props map[string]string
	// SQL is the raw SQL query.
	SQL string
	// SourceFile is the file where this query declared.
	SourceFile string
	// SourceLine is the line number where this query declared.
	SourceLine int
}
