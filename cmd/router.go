package main

import (
	"net/http"
)

type UserHandler interface {
	Login(http.ResponseWriter, *http.Request)
}

func initRoutes(h UserHandler) *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("GET /v1/api/login", h.Login)
	return m
}
