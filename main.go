package main

import (
	"gopher/infra/db/dbimpl"
	"gopher/services/product"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var ginMode string
	if os.Getenv("GIN_MODE") == "dev" {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)
	r := gin.Default()

	db := dbimpl.New(&product.Product{})

	ph := product.NewHandlers(db)
	r.Group("/products").GET("/", ph.GetProducts).GET("/:id", ph.FindById)

	log.Println("🚀 Server is running")
	r.Run()
}
