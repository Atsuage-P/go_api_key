package main

import (
	"log/slog"
	"os"

	"api_key_test/controller"
	"api_key_test/oapi"
	"api_key_test/structlog"

	"github.com/labstack/echo/v4"
)

func main() {
	slog.SetDefault(slog.New(structlog.NewLogHandler(slog.NewJSONHandler(os.Stdout, nil))))

	e := echo.New()
	api := &controller.APIController{}
	oapi.RegisterHandlers(e, api)
	e.Logger.Fatal(e.Start(":8080"))
}
