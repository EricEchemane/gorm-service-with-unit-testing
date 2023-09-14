package main

import (
	"context"
	"gopher/infra/db/dbimpl"
	"gopher/services/identity"
	"gopher/services/product"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	productServer := product.NewServer(&g, db, router)
	identityServer := identity.NewServer(&g, db, router)

	<-done

	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := productServer.Shutdown(ctx); err != nil {
		log.Fatal("Product error while shutting down")
		os.Exit(1)
	}
	if err := identityServer.Shutdown(ctx); err != nil {
		log.Fatal("Identity error while shutting down")
		os.Exit(1)
	}

	log.Println("All servers shutdown gracefully")
}
