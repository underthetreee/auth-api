package repository

import (
	"context"

	"github.com/underthetreee/auth-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewAuthRepository(client *mongo.Client) *AuthRepository {
	coll := client.Database("test").Collection("tokens")
	return &AuthRepository{
		client: client,
		coll:   coll,
	}
}

func (r *AuthRepository) Store(ctx context.Context, token model.HashedRefreshToken) error {
	_, err := r.coll.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
