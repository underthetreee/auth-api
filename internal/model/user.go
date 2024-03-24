package model

import "github.com/beevik/guid"

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type User struct {
	GUID               guid.Guid
	HashedRefreshToken string
}
