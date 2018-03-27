package sql

import (
	"database/sql"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/szokodiakos/r8m8/logger"
)

// Transaction interface
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Commit() error
	Rollback() error
}

type transaction struct {
	tx *sqlx.Tx
}

func (t *transaction) Commit() error {
	logger.Get().WithFields(logrus.Fields{
		"operation": "Transaction Commit",
	}).Info()
	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	logger.Get().WithFields(logrus.Fields{
		"operation": "Transaction Rollback",
	}).Info()
	return t.tx.Rollback()
}

func (t *transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	logger.Get().WithFields(logrus.Fields{
		"query":     formatQueryString(query),
		"operation": "Exec",
		"input":     spew.Sprint(args),
	}).Info()
	return t.tx.Exec(query, args...)
}

func (t *transaction) Select(dest interface{}, query string, args ...interface{}) error {
	err := t.tx.Select(dest, query, args...)
	logger.Get().WithFields(logrus.Fields{
		"query":     formatQueryString(query),
		"operation": "Select",
		"input":     spew.Sprint(args),
		"output":    spew.Sprint(dest),
	}).Info()
	return err
}

func (t *transaction) Get(dest interface{}, query string, args ...interface{}) error {
	err := t.tx.Get(dest, query, args...)
	logger.Get().WithFields(logrus.Fields{
		"query":     formatQueryString(query),
		"operation": "Get",
		"input":     spew.Sprint(args),
		"output":    spew.Sprint(dest),
	}).Info()
	return err
}

func formatQueryString(query string) string {
	return withoutTabs(withoutNewLines(query))
}

func withoutNewLines(input string) string {
	return strings.Replace(input, "\n", "", -1)
}

func withoutTabs(input string) string {
	return strings.Replace(input, "\t", "", -1)
}

// NewSQLTransaction factory
func NewSQLTransaction(tx *sqlx.Tx) Transaction {
	return &transaction{
		tx: tx,
	}
}
