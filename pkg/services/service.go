package services

import (
	"log"

	"github.com/marshacb/dreamdeeply/models"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

// GetAllProducts queries and returns all products
func (s Service) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product

	err := s.DB.Raw(getAllProducts).Scan(&products).Error
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving all products")
	}

	return &products, nil
}
