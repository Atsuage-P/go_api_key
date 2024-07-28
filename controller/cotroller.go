package controller

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"api_key_test/env"
	"api_key_test/oapi"
	"api_key_test/structlog"

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

	defer func() {
		if v := recover(); v != nil {
			errMsg := fmt.Errorf("Recovered from: %v", v)
			slog.ErrorContext(ctx, errMsg.Error(), "method", "DeleteNumber")
			if err := c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"message": http.StatusText(http.StatusInternalServerError),
				}); err != nil {
				slog.ErrorContext(ctx, err.Error(), "method", "DeleteNumber")
			}
		}
	}()

	if err := validateAPIKey(params.XAPIKEY); err != nil {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "DeleteNumber")

		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req, err := validateRequestBody(c)
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), "method", "DeleteNumber")

		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	num := *req.Num - 1
	res := oapi.NumberRes{
		Num: &num,
	}
	slog.InfoContext(ctx, "success", "method", "DeleteNumber")

	return c.JSON(http.StatusOK, res)
}

func (a *APIController) PostNumber(c echo.Context, params oapi.PostNumberParams) error {
	ctx := c.Request().Context()
	ctx = structlog.WithValue(ctx, "params", params)

	defer func() {
		if v := recover(); v != nil {
			errMsg := fmt.Errorf("Recovered from: %v", v)
			slog.ErrorContext(ctx, errMsg.Error(), "method", "PostNumber")
			if err := c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"message": http.StatusText(http.StatusInternalServerError),
				}); err != nil {
				slog.ErrorContext(ctx, err.Error(), "method", "PostNumber")
			}
		}
	}()

	if err := validateAPIKey(params.XAPIKEY); err != nil {
		slog.ErrorContext(ctx, "Wrong API Key", "method", "PostNumber")

		return c.JSON(http.StatusUnauthorized, map[string]string{"message": http.StatusText(http.StatusUnauthorized)})
	}

	req, err := validateRequestBody(c)
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), "method", "PostNumber")

		return c.JSON(http.StatusBadRequest, map[string]string{"message": http.StatusText(http.StatusBadRequest)})
	}

	num := *req.Num + 1
	res := oapi.NumberRes{
		Num: &num,
	}
	slog.InfoContext(ctx, "success", "method", "PostNumber")

	return c.JSON(http.StatusOK, res)
}

func validateAPIKey(reqKey string) error {
	KeyError := errors.New("Wrong API Key")
	cfg := env.LoadEnv()
	if reqKey != cfg.APIKey {
		return KeyError
	}

	return nil
}

func validateRequestBody(c echo.Context) (*oapi.NumberReq, error) {
	req := new(oapi.NumberReq)
	if err := c.Bind(req); err != nil {
		return nil, &RequestBindError{Message: "Request Bind Error"}
	}
	if req.Num == nil {
		return nil, &MissingRequestError{Message: "Missing Request Body"}
	}

	return req, nil
}

type RequestBindError struct {
	Message string
}

func (e *RequestBindError) Error() string {
	return e.Message
}

type MissingRequestError struct {
	Message string
}

func (e *MissingRequestError) Error() string {
	return e.Message
}
