package product

type ProductService interface {
	GetAllProducts() (*[]Product, error)
}

// List returns all known products
func List(ctx context.Context, p *ProductService) (*[]Product, error) {
	products, err := p.GetAllProducts()
	if err != nil {
		return nil, err
	}
	
	return &products, nil
}

// GetAllProducts queries and returns all products
func (s *Service) GetAllProducts() (*[]Product, error) {
	var products []Product{}

	err := s.DB.Raw(getAllProducts).Scan(&producs).Error
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving all products")
	}

	return &products, nil
}
