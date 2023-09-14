package product

import (
	"gopher/infra/db"
	"gopher/middleware"

	// "gopher/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group, db db.IDB, router *gin.Engine) *http.Server {
	handlers := NewHandlers(db)
	r := router.Group("/products").Use(middleware.Auth())
	{
		r.GET("/", handlers.GetProducts)
		r.GET("/:id", handlers.FindById)
	}

	productServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	g.Go(func() error {
		log.Default().Println("🚀 Product server listening on port 8080")
		return productServer.ListenAndServe()
	})

	return productServer
}
