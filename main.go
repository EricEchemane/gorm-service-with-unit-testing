package main

import (
	"gopher/services/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	ph := product.NewHandlers()
	r.Group("/products").GET("/", ph.GetProducts).GET("/:id", ph.FindById)

	log.Println("🚀 Server is running")
	r.Run()
}
