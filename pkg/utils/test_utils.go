package utils

import (
	"dreamdeeply/pkg/services"
	"log"
	"os"

	dreamdeeply "github.com/marshacb/dreamdeeply"

	"github.com/jinzhu/gorm"
)

// TODO: abstract this to a proper config package
func getDBConnectionString() string {
	return os.Getenv("TEST_MYSQL_CONNECTION_STRING")
}

// TestEnv contains everything we need to run tests
type TestEnv struct {
	Service *services.Service
}

// NewTestService creates a service for use in tests
// TODO: build out test http server and storage layer***
func NewTestService(db *gorm.DB, logger *log.Logger) *services.Service {
	return &services.Service{
		Logger: logger,
		DB:     db,
	}
}

// InitializeTest creates testenv used for integration tests
func InitializeTest() *TestEnv {
	var connectionString = getDBConnectionString()
	if connectionString == "" {
		connectionString = "root:thelastword@tcp(127.0.0.1:3306)/dream_deeply_test?charset=utf8mb4&parseTime=True"
	}

	log.Printf("connection string %s", connectionString)

	db, err := dreamdeeply.DB(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)

	err = dreamdeeply.MigrateUp(db.DB(), "file://../.../db/migrations")
	if err != nil {
		logger.Fatalf("error migrating up %s", err.Error())
	}

	service := NewTestService(db, logger)

	te := TestEnv{
		Service: service,
	}

	return &te
}

// TearDownTest migrates db down and closes db connection
func TearDownTest(te *TestEnv) {
	err := dreamdeeply.MigrateDown(te.Service.DB(), "file://../../db/migrations")
	if err != nil {
		te.Service.Logger.Fatalf("error migrating down %s".err.Error())
	}

	te.Service.DB.Close()
}
