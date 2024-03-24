package v1

import (
	"net/http"

	"github.com/beevik/guid"
	"github.com/underthetreee/auth-api/internal/auth"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	guidParam := r.URL.Query().Get("guid")

	parsedGUID, err := guid.ParseString(guidParam)
	if err != nil {
		http.Error(w, "invalid guid", http.StatusBadRequest)
	}

	tokenPair, err := auth.GenTokenPair(parsedGUID.String())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
