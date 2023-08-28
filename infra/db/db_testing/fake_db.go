package dbtesting

import (
	"gopher/infra/db"
)

type fakeDB struct {
	ExpectedError  error
	ExpectedResult interface{}
}

func NewFakeDB(expectedError error, expectedResult interface{}) db.IDB {
	return &fakeDB{
		ExpectedError:  expectedError,
		ExpectedResult: expectedResult,
	}
}

func (f *fakeDB) RawScan(sql string, dest interface{}, values ...interface{}) error {
	return f.ExpectedError
}
