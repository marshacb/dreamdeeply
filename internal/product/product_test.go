package product_test

import (
	"context"
	"testing"

	"github.com/marshacb/dreamdeeply/internal/product"
)

type MockProduct struct{}

func (m *MockProduct) GetAllProducts() (*[]product.Product, error) {
	products := []product.Product{
		Product{
			Name: "Ocean Wave",
		},
	}

	return products, nil
}

func TestListProduct(t *testing.T) {
	ctx := context.Background()
	mp := MockProduct{}

	products := List(ctx, &mp)

	if len(products) != 1 {
		t.Errorf("products should not be empty")
	}
}
