package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// HealthCheckModule 서버의 헬스체크를 위한 fx 모듈
var HealthCheckModule = fx.Module(
	"go-test-examples/healthcheck",
	fx.Invoke(func(e *echo.Echo) {
		e.GET("", func(c echo.Context) error {
			return c.String(http.StatusOK, "ok")
		})
	}),
)
