package dbgen

import "regexp"

var (
	rxQueryProps  = regexp.MustCompile(`(?i)([^\s:,]+)\s*:\s*([^\s:,]+)`)
	rxQueryParams = regexp.MustCompile(`(?i)([!:]{1,2})([\w_\d]+)`)
	rxTableName   = regexp.MustCompile(`(?i)CREATE\s+(?:TEMPORARY\s+)?TABLE\s+(?:IF\s+(?:NOT\s+)?EXISTS\s+)?(\S+)`)
)

const sqlEnableForeignKey = `SET FOREIGN_KEY_CHECKS = 1`
const sqlDisableForeignKey = `SET FOREIGN_KEY_CHECKS = 0`

const sqlCreateDdlCacheTable = `
CREATE TABLE IF NOT EXISTS ddl_cache (
	table_name VARCHAR(150) NOT NULL,
	sql_text   TEXT         NOT NULL,
	PRIMARY KEY (table_name)
)  CHARACTER SET utf8mb4`

const sqlGetDdlCache = `
SELECT sql_text FROM ddl_cache
WHERE table_name = ?`

const sqlInsertDdlCache = `
INSERT INTO ddl_cache (table_name, sql_text)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE sql_text = VALUES(sql_text)`
