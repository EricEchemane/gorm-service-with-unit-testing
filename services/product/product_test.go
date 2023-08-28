package product

import (
	"fmt"
	dbtesting "gopher/infra/db/db_testing"
	"testing"
)

func TestGetProducts(t *testing.T) {
	t.Run("service.GetProducts should return products", func(t *testing.T) {
		fakeDb := dbtesting.NewFakeDB(nil, []Product{})
		service := NewService(fakeDb)
		result, err := service.GetProducts(100)
		fmt.Println("result:", result)
		if err != nil {
			t.Fatalf("Expected to have no error but got %q", err)
		}
	})
}
