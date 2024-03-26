package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	GUID  string `json:"guid"`
	Token string `json:"refresh_token"`
}

type HashedRefreshToken struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	GUID  string             `bson:"guid"`
	Token string             `bson:"hashed_token"`
}
