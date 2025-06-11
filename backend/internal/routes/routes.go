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

func RegisterNoteRoutes(r *mux.Router) {
	noteRoutes := r.PathPrefix("/api/notes").Subrouter()

	// All note routes are protected with authentication middleware
	noteRoutes.HandleFunc("/", middleware.AuthMiddleware(handler.CreateNote)).Methods("POST")
	noteRoutes.HandleFunc("/", middleware.AuthMiddleware(handler.GetAllNotes)).Methods("GET")
	noteRoutes.HandleFunc("/{id}", middleware.AuthMiddleware(handler.GetNote)).Methods("GET")
	noteRoutes.HandleFunc("/{id}", middleware.AuthMiddleware(handler.DeleteNote)).Methods("DELETE")
	noteRoutes.HandleFunc("/{id}", middleware.AuthMiddleware(handler.UpdateNote)).Methods("PUT")
}
