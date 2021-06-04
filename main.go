package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RadhiFadlillah/dbgen/internal/sqlparser"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	// Connect to database
	db, err := openDB()
	checkError(err)
	defer db.Close()

	// Create parser
	ps := sqlparser.Parser{
		TmpDB:  db,
		SrcDir: "query",
	}

	err = ps.Parse()
	checkError(err)
}

func openDB() (*sqlx.DB, error) {
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
