package product

import (
	"context"

	"dreamdeeply/interfaces"

	"dreamdeeply/models"
)

// List returns all known products
func List(ctx context.Context, a interfaces.API) (*[]models.Product, error) {
	products, err := a.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
