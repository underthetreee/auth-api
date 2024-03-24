package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	v1 "github.com/underthetreee/auth-api/internal/transport/http/v1"
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
	_ = c

	userHandler := v1.NewUserHandler()
	router := initRoutes(userHandler)
	fmt.Printf("server is listening on %s\n", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
