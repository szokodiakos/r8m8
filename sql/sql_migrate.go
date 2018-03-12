package sql

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// Execute migrate
func Execute(db *sqlx.DB, dialect string) {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql/migrations",
	}

	n, err := migrate.Exec(db.DB, dialect, migrations, migrate.Up)

	if err != nil {
		log.Fatal("Error during migration: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
