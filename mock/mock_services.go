package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type UserService struct {
    UserFn func(id string) (*habitapi.User, error)
    UserInvoked bool

    UsersFn func() ([]*habitapi.User, error)
    UsersInvoked bool

    CreateUserFn func(dto habitapi.CreateUserDTO) (*habitapi.User, error)
    CreateUserInvoked bool

    DeleteUserFn func(id string) error
    DeleteUserInvoked bool
}

func (s *UserService) User(id string) (*habitapi.User, error) {
    s.UserInvoked = true
    return s.UserFn(id)
}

func (s *UserService) Users() ([]*habitapi.User, error) {
    s.UsersInvoked = true
    return s.UsersFn()
}

func (s *UserService) CreateUser(dto habitapi.CreateUserDTO) (*habitapi.User, error) {
    s.CreateUserInvoked = true
    return s.CreateUserFn(dto)
}

func (s *UserService) DeleteUser(id string) error {
    s.DeleteUserInvoked = true
    return s.DeleteUserFn(id)
}
