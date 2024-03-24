package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/underthetreee/auth-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret key")

func GenTokenPair(guid string) (*model.TokenPair, error) {
	token, err := genJWTToken(guid)
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

func genJWTToken(guid string) (string, error) {
	jwtTokenExpTime := time.Now().Add(time.Minute * 60).Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"guid": guid,
			"exp":  jwtTokenExpTime,
		})

	token, err := t.SignedString(secretKey)
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
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	return token, err
}

func hashRefreshToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}
