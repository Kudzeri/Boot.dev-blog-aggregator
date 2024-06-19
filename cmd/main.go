package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
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
