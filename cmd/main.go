package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/underthetreee/auth-api/internal/config"
	"github.com/underthetreee/auth-api/internal/repository"
	"github.com/underthetreee/auth-api/internal/service"
	api "github.com/underthetreee/auth-api/internal/transport/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		return fmt.Errorf("mongo: %w", err)
	}

	repo := repository.NewAuthRepository(cfg, c)
	svc := service.NewAuthService(cfg, repo)
	auth := api.NewAuthHandler(cfg, svc)
	router := initRoutes(auth)

	log.Printf("server is listening on %s\n", cfg.HTTP.Port)
	if err := http.ListenAndServe(cfg.HTTP.Port, router); err != nil {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}
