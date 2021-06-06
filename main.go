package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RadhiFadlillah/dbgen/internal/sqlparser"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	// Open temporary database
	tmpDB, err := openTmpDB()
	checkError(err)
	defer tmpDB.Close()

	// Parse SQL files
	ps := sqlparser.Parser{
		TmpDB:  tmpDB,
		SrcDir: "query",
	}

	_, _, _, err = ps.Parse()
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

func checkError(err error) {
	if err != nil && err != sql.ErrNoRows {
		logrus.Panicln(err)
	}
}
