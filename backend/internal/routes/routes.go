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
	// All note routes are protected with authentication middleware
	r.HandleFunc("/api/notes", middleware.AuthMiddleware(handler.CreateNote)).Methods("POST")
	r.HandleFunc("/api/notes", middleware.AuthMiddleware(handler.GetAllNotes)).Methods("GET")
	r.HandleFunc("/api/notes/{id}", middleware.AuthMiddleware(handler.GetNote)).Methods("GET")
	r.HandleFunc("/api/notes/{id}", middleware.AuthMiddleware(handler.DeleteNote)).Methods("DELETE")
	r.HandleFunc("/api/notes/{id}", middleware.AuthMiddleware(handler.UpdateNote)).Methods("PUT")
}
