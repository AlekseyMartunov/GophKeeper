package migration

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"path/filepath"
)

const homeDir = "GophKeeper"

func UserMigrationsUp(dsn string) error {
	path, err := getHomeDir()
	if err != nil {
		return err
	}

	path = filepath.Join("file:", path, "migrations", "users")

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

func UserMigrationsDown(dsn string) error {
	path, err := getHomeDir()
	if err != nil {
		return err
	}

	path = filepath.Join("file:", path, "migrations", "users")

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

func getHomeDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for filepath.Base(path) != homeDir {
		path = filepath.Dir(path)
	}
	return path, nil
}
