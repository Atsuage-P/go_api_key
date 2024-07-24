package main

import (
	"log"
	"net/http"

	"api_key_test/oapi"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api := &apiController{}
	oapi.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":8080"))
}

type apiController struct{}

// インターフェース実装の確認
var _ oapi.ServerInterface = (*apiController)(nil)

func (a *apiController) GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
}

func (a *apiController) DeleteNumber(ctx echo.Context, params oapi.DeleteNumberParams) error {
	cfg := LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	return ctx.JSON(http.StatusOK, (*req.Num)-1)
}

func (a *apiController) PostNumber(ctx echo.Context, params oapi.PostNumberParams) error {
	cfg := LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	return ctx.JSON(http.StatusOK, (*req.Num)+1)
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
