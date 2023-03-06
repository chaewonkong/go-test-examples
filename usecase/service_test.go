package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chaewonkong/go-test-examples/domain"
	"github.com/chaewonkong/go-test-examples/mocks"
	"github.com/chaewonkong/go-test-examples/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.nexon.com/infraleadingtech/go-modules/_test"
	"go.uber.org/fx"
)

func mockRepoFindAll(tb testing.TB, expected []domain.User, err error) domain.UserRepository {
	_repo := mocks.NewUserRepository(tb)
	_repo.EXPECT().FindAll().Return(expected, err)
	return _repo
}

// TestFindAll mockRepository를 직접 생성해 사용하는 패턴
func TestFindAll(t *testing.T) {
	t.Run("성공적으로 모든 유저를 반환", func(t *testing.T) {
		mockUsers := []domain.User{
			{Name: "Leon", Age: 31},
		}
		repo := mockRepoFindAll(t, mockUsers, nil)
		svc := usecase.NewUserService(repo)

		users, err := svc.FindAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, users, mockUsers)
	})

	t.Run("에러 발생 시 성공적으로 에러를 반환", func(t *testing.T) {
		testError := errors.New("test error")

		repo := mockRepoFindAll(t, []domain.User{}, testError)
		svc := usecase.NewUserService(repo)

		users, err := svc.FindAll(context.Background())
		assert.EqualError(t, err, testError.Error())
		assert.Empty(t, users)
	})
}

/*
fx활용

*/

// findOneNotFoundRepoModule mockRepository를 fx.Module로 만들어 주입하는 방법
// Not Found error를 발생시키는 repository를 모듈로 만들고 DI를 통해 주입해 테스트
var findOneNotFoundRepoModule = fx.Module(
	"usecase/mock_repo/find_one/not_found",
	fx.Provide(func(tb testing.TB) domain.UserRepository {
		_repo := mocks.NewUserRepository(tb)
		_repo.EXPECT().FindOne(mock.Anything).Return(
			domain.User{},
			domain.ErrNotFound,
		)
		return _repo
	}),
)

// TestFindOne_NotFound 사용자를 찾을 수 없는 경우 ErrNotFound 에러 발생
func TestFindOne_NotFound(t *testing.T) {
	f := func(svc usecase.UserService) {
		user, err := svc.FindOne(context.Background(), "존재하지 않는 사용자명")
		assert.EqualError(t, err, domain.ErrNotFound.Error())
		assert.Zero(t, user)
	}

	// _test.NewForTest는 uber-go/fx에서 testing시 사용되는 방법임.
	// https://github.com/uber-go/fx/blob/6285a021bb9591359c552ffd9031b8709f6cf604/app_test.go#L49
	_app := _test.NewForTest(
		t,
		fx.Invoke(f),
		fx.Provide(usecase.NewUserService),
		findOneNotFoundRepoModule,
	)
	_app.RequireStart().RequireStop()
}

// mockRepoFindOneModule repo를 모듈로 제공하는 함수
func mockRepoFindOneModule(tb testing.TB, expected domain.User, err error) fx.Option {
	return fx.Module(
		"usecase/mock_repo/find_one/success",
		fx.Provide(func(tb testing.TB) domain.UserRepository {
			_repo := mocks.NewUserRepository(tb)
			_repo.EXPECT().FindOne(mock.Anything).Return(expected, err)
			return _repo
		}),
	)
}
func TestFindOne_Success(t *testing.T) {
	mockUser := domain.User{
		Name: "Simon",
		Age:  26,
	}
	mockRepoModule := mockRepoFindOneModule(t, mockUser, nil)
	f := func(svc usecase.UserService) {
		user, err := svc.FindOne(context.Background(), "Leon")
		assert.NoError(t, err)
		assert.Equal(t, user, mockUser)
	}

	_app := _test.NewForTest(
		t,
		fx.Invoke(f),
		fx.Provide(usecase.NewUserService),
		mockRepoModule,
	)
	_app.RequireStart().RequireStop()

}
