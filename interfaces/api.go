package interfaces

import (
	"github.com/marshacb/dreamdeeply/models"
)

// API represents the method set of behavior of dreamdeeply
type API interface {
	GetAllProducts() (*[]models.Product, error)
}
