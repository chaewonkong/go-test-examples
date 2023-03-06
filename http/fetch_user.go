package http

import (
	"net/http"

	"github.com/chaewonkong/go-test-examples/domain"
	"github.com/chaewonkong/go-test-examples/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	// FetchUser 사용자를 가져오는 API 인터페이스
	FetchUser interface {
		FindAll(c echo.Context) error
		FindOne(c echo.Context) error
	}

	fetchUser struct {
		service usecase.UserService
	}
)

// NewFetchUser 사용자를 가져오는 API 생성자
func NewFetchUser(svc usecase.UserService) FetchUser {
	return fetchUser{svc}
}

// FetchUsers 모든 사용자를 조회하는 핸들러 함수
func (f fetchUser) FindAll(c echo.Context) error {
	users, err := f.service.FindAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// FindOne 사용자의 name으로 특정 사용자를 조회하는 핸들러 함수
func (f fetchUser) FindOne(c echo.Context) error {
	name := c.Param("name")
	user, err := f.service.FindOne(c.Request().Context(), name)

	if err != nil {
		if err == domain.ErrNotFound {

			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// FetchUserModule 사용자 조회에 사용되는 fx 모듈
var FetchUserModule = fx.Module(
	"go-test-examples/fetch_user",
	fx.Invoke(func(e *echo.Echo, f FetchUser) {
		u := e.Group("users")
		u.GET("", f.FindAll)
		u.GET("/:name", f.FindOne)
	}),
	fx.Provide(NewFetchUser),
)
