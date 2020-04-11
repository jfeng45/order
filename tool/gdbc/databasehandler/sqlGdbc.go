// Package handles low level database access including transaction through *sql.Tx or *sql.DB
package databasehandler

import (
	"database/sql"
)

// SqlDBTx is the concrete implementation of sqlGdbc by using *sql.DB
type SqlDBTx struct {
	DB *sql.DB
}

// SqlConnTx is the concrete implementation of sqlGdbc by using *sql.Tx
type SqlConnTx struct {
	DB *sql.Tx
}

func (sdt *SqlDBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdt.DB.Exec(query, args...)
}

func (sdt *SqlDBTx) Prepare(query string) (*sql.Stmt, error) {
	return sdt.DB.Prepare(query)
}

func (sdt *SqlDBTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdt.DB.Query(query, args...)
}

func (sdt *SqlDBTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return sdt.DB.QueryRow(query, args...)
}

func (sdb *SqlConnTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdb.DB.Exec(query, args...)
}

func (sdb *SqlConnTx) Prepare(query string) (*sql.Stmt, error) {
	return sdb.DB.Prepare(query)
}

func (sdb *SqlConnTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdb.DB.Query(query, args...)
}

func (sdb *SqlConnTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return sdb.DB.QueryRow(query, args...)
}

