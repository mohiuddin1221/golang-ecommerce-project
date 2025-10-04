package db

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// RunMigration runs all pending migrations
func RunMigration(dbUrl string) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create migrate driver: %v", err)
	}

	// Absolute path to migrations folder
	_, b, _, _ := runtime.Caller(0)
	migrationsPath := filepath.Join(filepath.Dir(b), "../../migrations")
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Migration instance creation failed: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully ✅")
}
