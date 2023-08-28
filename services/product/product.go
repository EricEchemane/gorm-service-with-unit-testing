package product

import "gopher/infra/db"

type service struct {
	db db.IDB
}

func NewService(db db.IDB) *service {
	return &service{db: db}
}

func (s *service) GetProducts(limit uint) ([]Product, error) {
	var products []Product
	err := s.db.RawScan("SELECT * FROM products LIMIT ?", &products, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
