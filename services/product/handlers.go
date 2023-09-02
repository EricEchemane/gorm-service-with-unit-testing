package product

import (
	"gopher/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandlers struct {
	s *Store
}

func NewHandlers(db db.IDB) productHandlers {
	return productHandlers{
		s: NewStore(db),
	}
}

func (h productHandlers) GetProducts(c *gin.Context) {
	products, err := h.s.GetProducts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func (h productHandlers) FindById(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := h.s.FindById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if product == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}
