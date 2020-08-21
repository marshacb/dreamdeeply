package product_test

import (
	"context"
	"testing"

	"dreamdeeply/models"
	"dreamdeeply/pkg/product"
)

type MockProduct struct{}

func (m *MockProduct) GetAllProducts() (*[]models.Product, error) {
	products := []models.Product{
		{
			Name: "Ocean Wave",
		},
	}

	return &products, nil
}

func TestListProduct(t *testing.T) {
	ctx := context.Background()
	mp := &MockProduct{}

	products, _ := product.List(ctx, mp)

	if len(*products) != 1 {
		t.Errorf("products should not be empty")
	}
}
