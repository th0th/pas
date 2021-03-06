package pas

import (
	"database/sql"
	"time"
)

type inserter interface {
	InsertSQL(timestamp time.Time) (sql string, values []interface{})
	CreateTableSQL() string
	ExistingColumns(*sql.DB) (map[string]struct{}, error)
	AlterTableSQL(existingColumns map[string]struct{}) string
}
