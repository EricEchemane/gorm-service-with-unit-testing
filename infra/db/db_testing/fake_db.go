package dbtesting

import (
	"gopher/infra/db"
	"reflect"
)

type Expectations struct {
	ExpectedError  error
	ExpectedResult interface{}
	TestName       string
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
	if f.ExpectedResult != nil && isAPointer(f.ExpectedResult) {
		destValue := reflect.ValueOf(dest)
		sourceValue := reflect.ValueOf(f.ExpectedResult)
		destValue.Elem().Set(sourceValue)
	}
	return f.ExpectedError
}

func (f *fakeDB) Insert(dest interface{}) error {
	if f.ExpectedResult != nil && isAPointer(f.ExpectedResult) {
		destValue := reflect.ValueOf(dest)
		sourceValue := reflect.ValueOf(f.ExpectedResult)
		destValue.Elem().Set(sourceValue.Elem())
	}
	return f.ExpectedError
}

func isAPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}
