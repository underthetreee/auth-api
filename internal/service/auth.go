package service

import (
	"context"

	"github.com/underthetreee/auth-api/internal/auth"
	"github.com/underthetreee/auth-api/internal/config"
	"github.com/underthetreee/auth-api/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository interface {
	StoreToken(context.Context, model.HashedRefreshToken) error
	DeleteToken(context.Context, model.RefreshToken) error
}

type AuthService struct {
	cfg  *config.Config
	repo AuthRepository
}

func NewAuthService(cfg *config.Config, repo AuthRepository) *AuthService {
	return &AuthService{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *AuthService) StoreRefreshToken(ctx context.Context, token model.RefreshToken) error {
	hashedToken, err := auth.HashRefreshToken(token.Token)
	if err != nil {
		return err
	}

	t := model.HashedRefreshToken{
		ID:    primitive.NewObjectID(),
		GUID:  token.GUID,
		Token: hashedToken,
	}

	if err := s.repo.StoreToken(ctx, t); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) RefreshAccessToken(ctx context.Context, token model.RefreshToken) (*model.TokenPair, error) {
	if err := s.repo.DeleteToken(ctx, token); err != nil {
		return nil, err
	}

	tokenPair, err := auth.GenTokenPair(token.GUID, s.cfg.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	t := model.RefreshToken{
		GUID:  token.GUID,
		Token: tokenPair.RefreshToken,
	}

	if err := s.StoreRefreshToken(ctx, t); err != nil {
		return nil, err
	}

	return tokenPair, nil
}
