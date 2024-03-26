package main

import (
	"net/http"
)

type AuthHandler interface {
	Auth(http.ResponseWriter, *http.Request)
	Refresh(http.ResponseWriter, *http.Request)
}

func initRoutes(h AuthHandler) http.Handler {
	m := http.NewServeMux()

	m.HandleFunc("POST /api/auth", h.Auth)
	m.HandleFunc("POST /api/refresh", h.Refresh)
	return m
}
