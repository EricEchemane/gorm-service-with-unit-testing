package identity

import (
	"gopher/infra/db"
)

type Store struct {
	db db.IDB
}

func NewStore(db db.IDB) *Store {
	return &Store{db: db}
}

func (s *Store) FindByUsername(username string) (*User, error) {
	var user *User
	err := s.db.RawScan("SELECT * FROM users WHERE username = ?", &user, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) Create(dto *CreateIdentityDTO) (*User, error) {
	user := User{
		Username: dto.Username,
		Password: dto.Password,
		Status:   Active,
		Role:     Normal,
	}
	err := s.db.Insert(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
