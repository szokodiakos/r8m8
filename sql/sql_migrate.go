package sql

import (
	"database/sql"
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

// Execute migrate
func Execute(db *sql.DB, dialect string) {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql/migrations",
	}

	n, err := migrate.Exec(db, dialect, migrations, migrate.Up)

	if err != nil {
		log.Fatal("Error during migration: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
