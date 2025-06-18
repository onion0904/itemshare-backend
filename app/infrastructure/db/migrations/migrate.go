package migrations

import (
	"database/sql"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var once sync.Once

func Migrate(db *sql.DB) {
	once.Do(func() {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file:///app/infrastructure/db/migrations",
			"postgres", driver,
		)
		if err != nil {
			panic(err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
	})
}
