package dreamdeeply

import (
	"github.com/jinzhu/gorm"
	// mysql imported locally to work with database/sql
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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
