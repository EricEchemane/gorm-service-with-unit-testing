package main

import (
	"gopher/infra/db/dbimpl"
	"gopher/services/product"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	var ginMode string
	if os.Getenv("GIN_MODE") == "dev" {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	db := dbimpl.New(&product.Product{})

	product.NewServer(&g, db)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
