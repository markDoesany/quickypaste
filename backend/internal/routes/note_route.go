package routes

import (
	"github.com/gorilla/mux"
	"github.com/smokeyghost/QuickyPaste/internal/handler"
	"github.com/smokeyghost/QuickyPaste/internal/middleware"
)

func RegisterNoteRoutes(r *mux.Router) {
	r.HandleFunc("/api/notes", middleware.AuthMiddleware(handler.CreateNote)).Methods("POST")
	r.HandleFunc("/api/notes", middleware.AuthMiddleware(handler.GetAllNotes)).Methods("GET")
	r.HandleFunc("/api/notes/{id}", handler.GetNote).Methods("GET")
	r.HandleFunc("/api/notes/{id}", middleware.AuthMiddleware(handler.DeleteNote)).Methods("DELETE")
	r.HandleFunc("/api/notes/{id}", middleware.AuthMiddleware(handler.UpdateNote)).Methods("PUT")
}
