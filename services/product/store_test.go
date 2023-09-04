package product_test

import (
	"errors"
	dbtesting "gopher/infra/db/db_testing"
	"gopher/services/product"
	"testing"
)

func TestGetProducts(t *testing.T) {
	testCases := []dbtesting.Expectations{
		{
			TestName:       "service.GetProducts should return products",
			ExpectedError:  nil,
			ExpectedResult: []product.Product{},
		},
		{
			TestName:       "service.GetProducts should return error",
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := product.NewStore(fakeDb)
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
			TestName:       "service.FindById should return products",
			ExpectedError:  nil,
			ExpectedResult: &product.Product{},
		},
		{
			TestName:       "service.FindById should return error",
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			fakeDb := dbtesting.NewFakeDB(tc)
			service := product.NewStore(fakeDb)
			_, err := service.FindById("1")
			if tc.ExpectedError == nil && err != nil {
				t.Fatalf("Expected to have no error but got %q", err)
			} else if tc.ExpectedError != nil && err == nil {
				t.Fatalf("Unexpected error: %q", err)
			}
		})
	}
}
