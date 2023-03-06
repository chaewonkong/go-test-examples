package domain

import (
	"errors"
	"fmt"
)

// ErrNotFound 데이터를 찾을 수 없는 경우 error
var ErrNotFound = errors.New("Not Found")

type (
	// UserRepository 사용자 repository
	UserRepository interface {
		FindAll() ([]User, error)
		FindOne(name string) (User, error)
	}

	userRepositry struct {
		users []User
	}
)

// NewUserRepository UserRepository 생성자
func NewUserRepository() UserRepository {
	users := []User{
		{Name: "Leon", Age: 31},
		{Name: "Jane", Age: 29},
	}
	return userRepositry{users}
}

// FindAll 모든 사용자를 조회하는 메서드
func (u userRepositry) FindAll() (result []User, err error) {
	result = u.users
	return
}

// FindOne 사용자의 name으로 특정 사용자를 조회하는 메서드
func (u userRepositry) FindOne(name string) (result User, err error) {
	for _, user := range u.users {
		if name == user.Name {
			result = user
			return
		}
	}
	fmt.Printf("user not found:%v", name)

	// user를 찾지 못한 경우
	err = ErrNotFound
	return
}
