package service

import (
	"context"

	"github.com/underthetreee/auth-api/internal/auth"
	"github.com/underthetreee/auth-api/internal/model"
)

type AuthRepository interface {
	Store(context.Context, model.HashedRefreshToken) error
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) StoreRefreshToken(ctx context.Context, token model.RefreshToken) error {
	hashedToken, err := auth.HashRefreshToken(token.Token)
	if err != nil {
		return err
	}

	t := model.HashedRefreshToken{
		GUID:  token.GUID,
		Token: hashedToken,
	}

	if err := s.repo.Store(ctx, t); err != nil {
		return err
	}
	return nil
}
