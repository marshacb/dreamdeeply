package utils

import (
	"dreamdeeply/pkg/services"
	"os"
)

// TODO: abstract this to a proper config package
func getDBConnectionString() string {
	return os.Getenv("TEST_MYSQL_CONNECTION_STRING")
}

// TestEnv contains everything we need to run tests
type TestEnv struct {
	Service *services.Service
}
