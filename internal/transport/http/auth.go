package transport

import (
	"context"
	"log"
	"net/http"

	"github.com/beevik/guid"
	"github.com/underthetreee/auth-api/internal/auth"
	"github.com/underthetreee/auth-api/internal/config"
	"github.com/underthetreee/auth-api/internal/model"
)

type AuthService interface {
	StoreRefreshToken(context.Context, model.RefreshToken) error
	RefreshAccessToken(context.Context, model.RefreshToken) (*model.TokenPair, error)
}

type AuthHandler struct {
	cfg *config.Config
	svc AuthService
}

func NewAuthHandler(cfg *config.Config, svc AuthService) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
		svc: svc,
	}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var (
		guidParam = r.FormValue("guid")
	)

	if !guid.IsGuid(guidParam) {
		http.Error(w, "invalid input", http.StatusBadRequest)
		log.Printf("invalid guid: %s", guidParam)
		return
	}

	tokenPair, err := auth.GenTokenPair(guidParam, h.cfg.JWT.SecretKey)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	t := model.RefreshToken{
		GUID:  guidParam,
		Token: tokenPair.RefreshToken,
	}

	if err = h.svc.StoreRefreshToken(r.Context(), t); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusOK, tokenPair)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var (
		guidParam  = r.FormValue("guid")
		tokenParam = r.FormValue("refresh_token")
	)

	if !guid.IsGuid(guidParam) {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if tokenParam == "" {
		http.Error(w, "invalid token", http.StatusBadRequest)
		log.Println("empty token")
		return
	}

	t := model.RefreshToken{
		GUID:  guidParam,
		Token: tokenParam,
	}

	tokenPair, err := h.svc.RefreshAccessToken(r.Context(), t)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	JSONResponse(w, http.StatusOK, tokenPair)
}
