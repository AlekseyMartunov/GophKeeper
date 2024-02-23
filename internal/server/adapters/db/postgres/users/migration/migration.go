package migration

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func UsersMigration(dsn string) error {
	_, path, _, _ := runtime.Caller(0)
	upMigrationPath := filepath.Dir(path)

	upMigrationPath = "file:" + upMigrationPath

	m, err := migrate.New(upMigrationPath, dsn)
	if err != nil {
		return fmt.Errorf("create migration instance error: %w", err)
	}

	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("user migration error: %w", err)
		}
	}

	return nil
}
