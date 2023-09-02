package identity

import "gopher/infra/db"

type Store struct {
	db db.IDB
}

func NewStore(db db.IDB) *Store {
	return &Store{db: db}
}
