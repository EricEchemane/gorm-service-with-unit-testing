package product

import (
	"errors"
	dbtesting "gopher/infra/db/db_testing"
	"testing"
)

func TestGetProducts(t *testing.T) {
	testCases := []dbtesting.Expectations{
		{
			ExpectedError:  nil,
			ExpectedResult: []Product{},
		},
		{
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("service.GetProducts should return products", func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := NewService(fakeDb)
			_, err := service.GetProducts()
			if tc.ExpectedError == nil && err != nil {
				t.Fatalf("Expected to have no error but got %q", err)
			} else if tc.ExpectedError != nil && err == nil {
				t.Fatalf("Unexpected error: %q", err)
			}
		})
	}

}

func TestFindById(t *testing.T) {
	testCases := []dbtesting.Expectations{
		{
			ExpectedError:  nil,
			ExpectedResult: &Product{},
		},
		{
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("service.FindById should return a product", func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := NewService(fakeDb)
			_, err := service.FindById("1")
			if tc.ExpectedError == nil && err != nil {
				t.Fatalf("Expected to have no error but got %q", err)
			} else if tc.ExpectedError != nil && err == nil {
				t.Fatalf("Unexpected error: %q", err)
			}
		})
	}

}
