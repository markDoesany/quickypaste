package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/smokeyghost/QuickyPaste/internal/config"
	"github.com/smokeyghost/QuickyPaste/internal/database"
	"github.com/smokeyghost/QuickyPaste/internal/routes"
)

func main() {
	config.LoadConfig()
	database.InitializeDB()
	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r)
	routes.RegisterNoteRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://quickypaste.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Printf("Server running  on Port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler))
}
