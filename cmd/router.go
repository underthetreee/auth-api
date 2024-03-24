package main

import (
	"net/http"
)

type UserHandler interface {
	Auth(http.ResponseWriter, *http.Request)
}

func initRoutes(h UserHandler) http.Handler {
	m := http.NewServeMux()

	m.HandleFunc("POST /api/auth", h.Auth)

	return m
}
