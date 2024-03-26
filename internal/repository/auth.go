package repository

import (
	"context"
	"errors"

	"github.com/underthetreee/auth-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (r *AuthRepository) StoreToken(ctx context.Context, token model.HashedRefreshToken) error {
	_, err := r.coll.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) DeleteToken(ctx context.Context, token model.RefreshToken) error {
	cur, err := r.coll.Find(ctx, bson.M{"guid": token.GUID})
	if err != nil {
		return err
	}
	var (
		hashed model.HashedRefreshToken
		found  bool
	)
	for cur.Next(ctx) {
		if err := cur.Decode(&hashed); err != nil {
			return err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(hashed.Token), []byte(token.Token)); err == nil {
			if _, err := r.coll.DeleteOne(ctx, bson.M{"_id": hashed.ID}); err != nil {
				return err
			}
			found = true
			break
		}
	}
	if !found {
		return errors.New("token not found")
	}
	return nil
}
