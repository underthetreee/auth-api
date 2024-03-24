package model

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type HashedRefreshToken struct {
	GUID  string `bson:"guid"`
	Token string `bson:"hashed_token"`
}

type RefreshToken struct {
	GUID  string
	Token string
}
