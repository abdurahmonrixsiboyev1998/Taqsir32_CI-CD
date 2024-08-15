package server

import (
	"90HM/internal/handlers"
	"net/http"
)

func New() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/", handlers.RootHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
