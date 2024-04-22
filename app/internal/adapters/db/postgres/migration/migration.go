package migration

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type config interface {
	MigrationPath() string
	PostgresDSN() string
}

func MigrationsUp(cfg config) error {
	path := cfg.MigrationPath()
	dsn := cfg.PostgresDSN()
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	path = filepath.Join("file:", workDir, path)
	m, err := migrate.New(path, dsn)
	if err != nil {
		return fmt.Errorf("create migration instance error: %w", err)
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("user migration error: %w", err)
		}
	}

	return nil
}

func MigrationsDown(cfg config) error {
	path := cfg.MigrationPath()
	dsn := cfg.PostgresDSN()

	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	path = filepath.Join("file:", workDir, path)

	m, err := migrate.New(path, dsn)
	if err != nil {
		return fmt.Errorf("create migration instance error: %w", err)
	}

	err = m.Down()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("user migration error: %w", err)
		}
	}
	return nil
}
