package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"api_key_test/oapi"
	"api_key_test/structlog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	slog.SetDefault(slog.New(structlog.NewLogHandler(slog.NewJSONHandler(os.Stdout, nil))))

	e := echo.New()
	api := &apiController{}
	oapi.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":8080"))
}

type apiController struct{}

// インターフェース実装の確認
var _ oapi.ServerInterface = (*apiController)(nil)

func (a *apiController) GetHello(c echo.Context) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "test", "hoge")
	slog.InfoContext(ctx, "success", "method", "GetHello")

	return c.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
}

func (a *apiController) DeleteNumber(c echo.Context, params oapi.DeleteNumberParams) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "params", params)

	cfg := LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "DeleteNumber")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := c.Bind(req); err != nil {
		slog.ErrorContext(ctx, "Request Bind Error", "method", "DeleteNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	slog.InfoContext(ctx, "success", "method", "DeleteNumber")
	return c.JSON(http.StatusOK, (*req.Num)-1)
}

func (a *apiController) PostNumber(c echo.Context, params oapi.PostNumberParams) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "params", params)

	cfg := LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "PostNumber")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := c.Bind(req); err != nil {
		slog.ErrorContext(ctx, "Request Bind Error", "method", "PostNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	slog.InfoContext(ctx, "success", "method", "PostNumber")
	return c.JSON(http.StatusOK, (*req.Num)+1)
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
