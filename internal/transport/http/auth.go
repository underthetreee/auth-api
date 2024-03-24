package transport

import (
	"context"
	"net/http"

	"github.com/beevik/guid"
	"github.com/underthetreee/auth-api/internal/auth"
	"github.com/underthetreee/auth-api/internal/model"
)

type AuthService interface {
	StoreRefreshToken(ctx context.Context, token model.RefreshToken) error
}

type AuthHandler struct {
	svc AuthService
}

func NewAuthHandler(svc AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	guidParam := r.URL.Query().Get("guid")

	parsedGUID, err := guid.ParseString(guidParam)
	if err != nil {
		http.Error(w, "invalid guid", http.StatusBadRequest)
		return
	}

	tokenPair, err := auth.GenTokenPair(parsedGUID.String())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	t := model.RefreshToken{
		GUID:  parsedGUID.String(),
		Token: tokenPair.RefreshToken,
	}
	if err = h.svc.StoreRefreshToken(context.Background(), t); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return

	}
	JSONResponse(w, http.StatusOK, tokenPair)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

}
