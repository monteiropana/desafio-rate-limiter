package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rate-limiter-desafio/internal/middleware"
	"github.com/rate-limiter-desafio/internal/repository/redisRepository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisClient := redisRepository.Config()
	redisRepo := redisRepository.NewRedisRepository(redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	wrappedMux := middleware.RateLimitMiddleware(redisRepo)(mux)

	http.ListenAndServe(":8080", wrappedMux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
