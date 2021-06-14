package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/RadhiFadlillah/dbgen"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Open temporary database
	tmpDB, err := openTmpDB()
	checkError(err)
	defer tmpDB.Close()

	// Parse SQL files
	ps := dbgen.SqlParser{
		TmpDB:  tmpDB,
		SrcDir: "query",
	}

	ddlQueries, selectQueries, execQueries, err := ps.Parse()
	checkError(err)

	// Generate code
	gen := dbgen.Generator{
		DstDir:              "storage",
		PackageName:         "storage",
		DdlQueries:          ddlQueries,
		SelectQueries:       selectQueries,
		ExecQueries:         execQueries,
		ColumnTypeConverter: columnTypeConverter,
		AdditionalImports: []string{
			"time",
			"encoding/json",
			"gopkg.in/guregu/null.v4",
			"github.com/shopspring/decimal",
		},
	}

	err = gen.Run()
	checkError(err)
}

func openTmpDB() (*sqlx.DB, error) {
	// Connect to database
	dataSource := fmt.Sprintf("%s:@tcp(%s)/%s", "radhi", "127.0.0.1:3306", "training")
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	db.SetConnMaxLifetime(time.Minute)

	return db, nil
}

func columnTypeConverter(column dbgen.Column) string {
	dbType := strings.ToUpper(column.DbType)

	if column.Nullable {
		switch dbType {
		case "INT", "BIGINT":
			return "null.Int"
		case "BINARY", "VARBINARY", "VARCHAR":
			return "null.String"
		case "DATE", "TIME", "DATETIME", "TIMESTAMP":
			return "mysql.NullTime"
		case "DECIMAL":
			return "decimal.NullDecimal"
		case "TEXT":
			return "json.RawMessage"
		case "TINYINT":
			return "null.Bool"
		}
	} else {
		switch dbType {
		case "INT", "BIGINT":
			return "int"
		case "BINARY", "VARBINARY", "VARCHAR":
			return "string"
		case "DATE", "TIME", "DATETIME", "TIMESTAMP":
			return "time.Time"
		case "DECIMAL":
			return "decimal.Decimal"
		case "TEXT":
			return "json.RawMessage"
		case "TINYINT":
			return "bool"
		}
	}

	return column.ScanType
}

func checkError(err error) {
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
}
