package main

import (
	"gopher/infra/db/dbimpl"
	"gopher/services/identity"
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
	router := gin.Default()

	db := dbimpl.New(&product.Product{}, &identity.User{})

	product.NewServer(&g, db, router)
	identity.NewServer(&g, db, router)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
