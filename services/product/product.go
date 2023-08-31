package product

import "gopher/infra/db"

type Service struct {
	db db.IDB
}

func NewService(db db.IDB) *Service {
	return &Service{db: db}
}

func (s *Service) GetProducts() ([]Product, error) {
	var products []Product
	err := s.db.RawScan("SELECT * FROM products LIMIT 10", &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) FindById(id string) (*Product, error) {
	var product *Product
	err := s.db.RawScan("SELECT * FROM products WHERE id = ?", &product, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
