package identity

import (
	"fmt"
	"gopher/infra/db"
	"gopher/infra/session"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	dto.Password, err = HashPassword(dto.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user, err := h.s.Create(&dto)

	if err != nil {
		if strings.Contains(err.Error(), "users_username_key") {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Username %s is not available", dto.Username)})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = ""
	c.IndentedJSON(http.StatusCreated, user)
}

func (h *identityHandlers) Login(c *gin.Context) {
	var dto LoginDTO
	c.Bind(&dto)

	user, err := h.s.FindByUsername(dto.Username)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	matched := CheckPasswordHash(dto.Password, user.Password)
	if !matched {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	user.Password = ""

	uuid := uuid.New()
	session_id := uuid.String()
	session.Set(session_id, user.Username)
	c.SetCookie("session_id", session_id, 60, "/", "localhost", true, true)
	c.JSON(http.StatusOK, "ok")
}
