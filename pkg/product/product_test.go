package product_test

import (
	"context"
	"testing"

	"github.com/marshacb/dreamdeeply/models"
)

type MockProduct struct{}

func (m *MockProduct) GetAllProducts() (*[]models.Product, error) {
	products := []product.Product{
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
