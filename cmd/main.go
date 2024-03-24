package main

import (
	"context"
	"log"
	"net/http"

	"github.com/underthetreee/auth-api/internal/repository"
	"github.com/underthetreee/auth-api/internal/service"
	api "github.com/underthetreee/auth-api/internal/transport/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	listenAddr = ":8080"
	mongoURI   = "mongodb://localhost:27017"
)

func main() {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoURI)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewAuthRepository(c)
	svc := service.NewAuthService(repo)
	authHandler := api.NewAuthHandler(svc)
	router := initRoutes(authHandler)

	log.Printf("server is listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
