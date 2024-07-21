package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func main() {
	helloHandler := http.HandlerFunc(hello)
	http.HandleFunc("/hello", checkAPIKeyMiddleware(helloHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "200 OK")
}

func checkAPIKeyMiddleware(next http.Handler) http.HandlerFunc {
	cfg := LoadEnv()
	apiKey := cfg.APIKey

	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key != apiKey {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Config struct {
	APIKey string `env:"API_KEY,notEmpty"`
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Load Env Error: %v", err)
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Env Parse Error: %v", err)
	}
	return &cfg
}
