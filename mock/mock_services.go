package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type UserService struct {
    GetByIdFn func(id string) (*habitapi.User, error)
    GetByIdInvoked bool

    ListFn func() ([]*habitapi.User, error)
    ListInvoked bool

    CreateFn func(dto habitapi.CreateUserDTO) (*habitapi.User, error)
    CreateInvoked bool

    DeleteFn func(id string) error
    DeleteInvoked bool
}

func (s *UserService) GetById(id string) (*habitapi.User, error) {
    s.GetByIdInvoked = true
    return s.GetByIdFn(id)
}

func (s *UserService) List() ([]*habitapi.User, error) {
    s.ListInvoked = true
    return s.ListFn()
}

func (s *UserService) Create(dto habitapi.CreateUserDTO) (*habitapi.User, error) {
    s.CreateInvoked = true
    return s.CreateFn(dto)
}

func (s *UserService) Delete(id string) error {
    s.DeleteInvoked = true
    return s.DeleteFn(id)
}
