package usecase

import (
	"context"

	"github.com/chaewonkong/go-test-examples/domain"
)

type (
	// UserService 사용자를 조회하는 서비스
	UserService interface {
		FindAll(c context.Context) ([]domain.User, error)
		FindOne(c context.Context, name string) (domain.User, error)
	}

	userService struct {
		repo domain.UserRepository
	}
)

// NewUserService 사용자를 조회하는 서비스 생성자
func NewUserService(r domain.UserRepository) UserService {
	return userService{r}
}

// FindAll 사용자 전체를 조회하는 메서드
func (u userService) FindAll(c context.Context) (users []domain.User, err error) {
	users, err = u.repo.FindAll()
	return
}

// FindOne 사용자의 name으로 특정 사용자를 조회하는 메서드
func (u userService) FindOne(c context.Context, name string) (user domain.User, err error) {
	user, err = u.repo.FindOne(name)
	return
}
