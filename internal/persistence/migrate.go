package persistence

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // used for migrations Path
	_ "github.com/lib/pq"                                // used for Migrate to work
	"github.com/rs/zerolog/log"
)

// Migrate upgrades a database using a directory of sql migrations.
func Migrate(db *sql.DB, path string, dbName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://"+path, dbName, driver)
	if err != nil {
		return err
	}

	if err := migrations.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msgf("no migrations to run: %s", err.Error())
			return nil
		}

		return err
	}

	return nil
}
