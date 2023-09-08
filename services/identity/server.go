package identity

import (
	"gopher/infra/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group, db db.IDB, router *gin.Engine) {
	handlers := NewHandlers(db)
	r := router.Group("/identity")
	{
		r.POST("/", handlers.CreateIdentity)
		r.POST("/login", handlers.Login)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	g.Go(func() error {
		log.Default().Println("🚀 Identity server listening on port 8080")
		return server.ListenAndServe()
	})
}
