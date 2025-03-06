package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/smokeyghost/QuickyPaste/internal/handler"
	"github.com/smokeyghost/QuickyPaste/internal/middleware"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/register", handler.Register).Methods("POST")
	r.HandleFunc("/api/login", handler.Login).Methods("POST")

	r.HandleFunc("/protected", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the protected route!"))
	})).Methods("GET")
}
