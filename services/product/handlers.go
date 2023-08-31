package product

import (
	"gopher/infra/db/dbimpl"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandlers struct {
	s *Service
}

func NewHandlers() productHandlers {
	db := dbimpl.New(&Product{})
	return productHandlers{
		s: NewService(db),
	}
}

func (h productHandlers) GetProducts(c *gin.Context) {
	products, err := h.s.GetProducts()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func (h productHandlers) FindById(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := h.s.FindById(id)
	if err != nil {
		return
	}
	if product == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}
