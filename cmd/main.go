package main

import (
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/iternal"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router.Use(c.Handler)

	V1Router := chi.NewRouter()
	V1Router.Get("/healthz", iternal.HandlerReadiness)
	V1Router.Get("/err", iternal.HandlerErr)
	router.Mount("/v1", V1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	fmt.Println("Server starting on port:" + portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
