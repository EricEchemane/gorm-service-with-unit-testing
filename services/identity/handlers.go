package identity

import (
	"fmt"
	"gopher/infra/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type identityHandlers struct {
	s *Store
}

func NewHandlers(db db.IDB) *identityHandlers {
	return &identityHandlers{
		s: NewStore(db),
	}
}

func (h *identityHandlers) CreateIdentity(c *gin.Context) {
	var dto CreateIdentityDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.s.Create(&dto)
	if strings.Contains(err.Error(), "users_username_key") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Username %s is not available", dto.Username)})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
