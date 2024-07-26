package controller

import (
	"api_key_test/env"
	"api_key_test/oapi"
	"api_key_test/structlog"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIController struct{}

// インターフェース実装の確認
var _ oapi.ServerInterface = (*APIController)(nil)

func (a *APIController) GetHello(c echo.Context) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "test", "hoge")
	slog.InfoContext(ctx, "success", "method", "GetHello")

	return c.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
}

func (a *APIController) DeleteNumber(c echo.Context, params oapi.DeleteNumberParams) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "params", params)

	cfg := env.LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "DeleteNumber")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := c.Bind(req); err != nil {
		slog.ErrorContext(ctx, "Request Bind Error", "method", "DeleteNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	if req.Num == nil {
		slog.ErrorContext(ctx, "Wrong Request Body", "method", "PostNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	slog.InfoContext(ctx, "success", "method", "DeleteNumber")
	return c.JSON(http.StatusOK, map[string]int{"num": (*req.Num) - 1})
}

func (a *APIController) PostNumber(c echo.Context, params oapi.PostNumberParams) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "params", params)

	cfg := env.LoadEnv()
	if params.XAPIKEY != cfg.APIKey {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "PostNumber")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req := new(oapi.NumberReq)
	if err := c.Bind(req); err != nil {
		slog.ErrorContext(ctx, "Request Bind Error", "method", "PostNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	if req.Num == nil {
		slog.ErrorContext(ctx, "Wrong Request Body", "method", "PostNumber")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	slog.InfoContext(ctx, "success", "method", "PostNumber")
	return c.JSON(http.StatusOK, map[string]int{"num": (*req.Num) + 1})
}
