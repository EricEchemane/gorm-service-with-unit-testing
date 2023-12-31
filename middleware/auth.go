package middleware

import (
	"gopher/infra/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session_id, err := c.Cookie("session_id")
		if err != nil || session_id == "" {
			log.Default().Println("session_id: ", session_id)
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		username := session.Get(session_id)
		if username == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
