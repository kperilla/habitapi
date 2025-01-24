package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type DeedService struct {
    GetByIdFn func(id string) (*habitapi.Deed, error)
    GetByIdInvoked bool

    ListFn func() ([]*habitapi.Deed, error)
    ListInvoked bool

    CreateFn func(dto habitapi.CreateDeedDTO) (*habitapi.Deed, error)
    CreateInvoked bool

    UpdateFn func(id string, dto habitapi.UpdateDeedDTO) (*habitapi.Deed, error)
    UpdateInvoked bool

    DeleteFn func(id string) error
    DeleteInvoked bool
}

func (s *DeedService) GetById(id string) (*habitapi.Deed, error) {
    s.GetByIdInvoked = true
    return s.GetByIdFn(id)
}

func (s *DeedService) List() ([]*habitapi.Deed, error) {
    s.ListInvoked = true
    return s.ListFn()
}

func (s *DeedService) Create(dto habitapi.CreateDeedDTO) (*habitapi.Deed, error) {
    s.CreateInvoked = true
    return s.CreateFn(dto)
}

func (s *DeedService) Update(id string, dto habitapi.UpdateDeedDTO) (*habitapi.Deed, error) {
    s.UpdateInvoked = true
    return s.UpdateFn(id, dto)
}

func (s *DeedService) Delete(id string) error {
    s.DeleteInvoked = true
    return s.DeleteFn(id)
}
