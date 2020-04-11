// package gdbc is created to represents low level database interfaces in order to have an unified way to
// access database handler.
// It is created to make it easier to handle certain database operations like transactions, database factory.
// It is ony a POC, not a mature solution
package gdbc

import (
	"database/sql"
)

// SqlGdbc (SQL Go database connection) is a wrapper for SQL database handler ( can be *sql.DB or *sql.Tx)
// It should be able to work with all SQL data that follows SQL standard.
type SqlGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
