package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/underthetreee/auth-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func GenTokenPair(guid string, secretKey string) (*model.TokenPair, error) {
	token, err := genJWTToken(guid, secretKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := genRefreshToken()
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func genJWTToken(guid string, secretKey string) (string, error) {
	jwtTokenExpTime := time.Now().Add(time.Hour * 1).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"guid": guid,
			"exp":  jwtTokenExpTime,
		})

	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func genRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, nil
}

func HashRefreshToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}
