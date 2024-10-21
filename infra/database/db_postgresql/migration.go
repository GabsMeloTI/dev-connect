package db_postgresql

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
	"strings"
)

func runMigrations(conn *sql.DB, environment string) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		errConnection(environment, err)
		return err
	}

	pwd, err := os.Getwd() // Verificar se há erro ao obter o diretório
	if err != nil {
		errConnection(environment, err)
		return err
	}

	migrationsPath := filepath.Join(pwd, "db/migration")

	// Garantir que o caminho tenha barras normais
	migrationsPath = strings.ReplaceAll(migrationsPath, "\\", "/")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath, // Corrigindo a formatação da URL
		"postgres", driver)
	if err != nil {
		errConnection(environment, err)
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		errConnection(environment, err)
		return err
	}

	return nil
}
