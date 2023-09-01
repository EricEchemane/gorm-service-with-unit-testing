package product_test

import (
	"errors"
	dbtesting "gopher/infra/db/db_testing"
	"gopher/services/product"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tj/assert"
)

func TestGetProductsHandler(t *testing.T) {
	testCases := []struct {
		TestName       string
		ExpectedError  error
		ExpectedStatus int
	}{
		{
			TestName:       "/products should return 200 status",
			ExpectedError:  nil,
			ExpectedStatus: 200,
		},
		{
			TestName:       "/products should return 400",
			ExpectedError:  errors.New("Some error"),
			ExpectedStatus: 500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()

			db := dbtesting.NewFakeDB(dbtesting.Expectations{
				ExpectedError: tc.ExpectedError,
			})
			handlers := product.NewHandlers(db)

			r.GET("/products", handlers.GetProducts)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/products", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.ExpectedStatus, w.Code)
		})

	}
}

func TestFindByIdHandler(t *testing.T) {
	testCases := []struct {
		TestName       string
		ExpectedError  error
		ExpectedResult *product.Product
		ExpectedStatus int
	}{
		{
			TestName:       "/products/:id should return 200 status",
			ExpectedError:  nil,
			ExpectedResult: &product.Product{},
			ExpectedStatus: 200,
		},
		{
			TestName:       "/products/:id should return 400",
			ExpectedError:  errors.New("Some error"),
			ExpectedResult: nil,
			ExpectedStatus: 400,
		},
		{
			TestName:       "/products/:id should return 404",
			ExpectedError:  nil,
			ExpectedResult: nil,
			ExpectedStatus: 404,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()

			db := dbtesting.NewFakeDB(dbtesting.Expectations{
				ExpectedError:  tc.ExpectedError,
				ExpectedResult: tc.ExpectedResult,
			})
			handlers := product.NewHandlers(db)

			r.GET("/products/1", handlers.FindById)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/products/1", nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.ExpectedStatus, w.Code)
		})
	}
}
