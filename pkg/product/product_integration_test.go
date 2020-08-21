package product_test

import (
	"dreamdeeply/pkg/utils"
	"log"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	te := utils.InitializeTest()
	defer utils.TearDownTest(te)

	log.Println("te is: ", te)
}

// func TestListProducts(t *testing.T) {}
