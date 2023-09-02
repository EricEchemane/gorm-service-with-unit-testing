package product

import "gopher/infra/db"

type Store struct {
	db db.IDB
}

func NewStore(db db.IDB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]Product, error) {
	var products []Product
	err := s.db.RawScan("SELECT * FROM products LIMIT 10", &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Store) FindById(id string) (*Product, error) {
	var product *Product
	err := s.db.RawScan("SELECT * FROM products WHERE id = ?", &product, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
