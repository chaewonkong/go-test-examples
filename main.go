package main

import (
	"github.com/chaewonkong/go-test-examples/domain"
	"github.com/chaewonkong/go-test-examples/http"
	"github.com/chaewonkong/go-test-examples/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		http.HealthCheckModule,
		http.FetchUserModule,
		fx.Provide(echo.New, domain.NewUserRepository, usecase.NewUserService),
		fx.Invoke(http.Run),
	).Run()
}
