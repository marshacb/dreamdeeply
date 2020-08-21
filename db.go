package dreamdeeply

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"

	// mysql imported locally to work with database/sql
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DB gets a connection to a database
func DB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	db.DB().Ping()
	db.DB().SetMaxOpenConns(25)
	db.DB().SetMaxIdleConns(25)

	return db, nil
}

// getMigrator creates a migrator from a connection
func getMigrator(db *sql.DB, migrationPath string) (*migrate.Migrate, error) {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	return migrate.NewWithDatabaseInstance(
		// TODO: How can we ensute the working dir is always the same so we can path to this migration folder?
		migrationPath,
		"mysql",
		driver,
	)
}

// MigrateUp creates a migrator from a connection
func MigrateUp(db *sql.DB, migrationPath string) error {
	m, err := getMigrator(db, migrationPath)
	if err != nil {
		log.Errorf(context.Background(), "error while migrating %s", err.Error())
		return err
	}

	m.Up()

	return nil
}

// MigrateDown brings the database to the initial state
// Note: this is only intended for test dbs, gated to ensure prd dbs aren't Down'd
func MigrateDown(db *sql.DB, migrationPath string) error {
	m, err := getMigrator(db, migrationPath)
	if err != nil {
		log.Errorf(context.Background(), "error while migrating %s", err.Error())
		return err
	}

	m.Down()

	return nil
}
