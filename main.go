package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			cfg := LoadEnv()
			apiKey := cfg.APIKey
			return key == apiKey, nil
		},
	}))

	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
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
