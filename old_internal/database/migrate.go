package database

import (
    "database/sql"
    "fmt"
    "log"
	"path/filepath"

    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(dsn string, path string) error {
    // создаём *sql.DB
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return fmt.Errorf("cannot open sql db: %w", err)
    }
    defer db.Close()

    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("cannot create migration driver: %w", err)
    }

	absPath, err := filepath.Abs(path)
    if err != nil {
        return fmt.Errorf("cannot get absolute path: %w", err)
    }

    mig, err := migrate.NewWithDatabaseInstance(
        "file://"+absPath,
        "postgres",
        driver,
    )
    if err != nil {
        return fmt.Errorf("cannot init migrate: %w", err)
    }

    err = mig.Up()
    if err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("migration failed: %w", err)
    }

    log.Println("Migrations applied!")
    return nil
}
