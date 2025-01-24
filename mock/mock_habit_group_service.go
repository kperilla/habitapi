package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type HabitGroupService struct {
    GetByIdFn func(id string) (*habitapi.HabitGroup, error)
    GetByIdInvoked bool

    ListFn func() ([]*habitapi.HabitGroup, error)
    ListInvoked bool

    CreateFn func(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error)
    CreateInvoked bool

    UpdateFn func(id string, dto habitapi.UpdateHabitGroupDTO) (*habitapi.HabitGroup, error)
    UpdateInvoked bool

    DeleteFn func(id string) error
    DeleteInvoked bool
}

func (s *HabitGroupService) GetById(id string) (*habitapi.HabitGroup, error) {
    s.GetByIdInvoked = true
    return s.GetByIdFn(id)
}

func (s *HabitGroupService) List() ([]*habitapi.HabitGroup, error) {
    s.ListInvoked = true
    return s.ListFn()
}

func (s *HabitGroupService) Create(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error) {
    s.CreateInvoked = true
    return s.CreateFn(dto)
}

func (s *HabitGroupService) Update(id string, dto habitapi.UpdateHabitGroupDTO) (*habitapi.HabitGroup, error) {
    s.UpdateInvoked = true
    return s.UpdateFn(id, dto)
}

func (s *HabitGroupService) Delete(id string) error {
    s.DeleteInvoked = true
    return s.DeleteFn(id)
}
