package services

import (
	"log"

	"github.com/jinzhu/gorm"
)

const getAllProducts = `
select * from products
`

// Service defines the services of the application
type Service struct {
	DB     *gorm.DB
	Logger *log.Logger
}

// New creates and returns a new service
func New(db *gorm.DB, logger *log.Logger) *Service {

	return &Service{
		DB:     db,
		Logger: logger,
	}
}
