package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/szokodiakos/r8m8/logger"
)

// Execute migrate
func Execute(db *sqlx.DB, dialect string) {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql/migrations",
	}

	n, err := migrate.Exec(db.DB, dialect, migrations, migrate.Up)

	if err != nil {
		logger.Get().Fatal("Error during migration", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
