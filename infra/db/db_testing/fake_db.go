package dbtesting

import (
	"gopher/infra/db"
)

type Expectations struct {
	ExpectedError  error
	ExpectedResult interface{}
}

type fakeDB struct {
	Expectations
}

func NewFakeDB(expectations Expectations) db.IDB {
	return &fakeDB{
		Expectations: expectations,
	}
}

func (f *fakeDB) RawScan(sql string, dest interface{}, values ...interface{}) error {
	return f.ExpectedError
}
