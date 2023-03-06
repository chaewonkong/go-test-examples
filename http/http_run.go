package http

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// Run HTTP 서버를 실행하는 함수
func Run(
	lifecycle fx.Lifecycle,
	e *echo.Echo,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// hooks는 blocking으로 동작하므로 separate goroutine으로 실행 필요
			// https://github.com/uber-go/fx/issues/627#issuecomment-399235227
			go func() {
				if err := e.Start(":8080"); err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
