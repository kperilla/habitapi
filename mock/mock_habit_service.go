package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type HabitService struct {
    GetByIdFn func(id string) (*habitapi.Habit, error)
    GetByIdInvoked bool

    ListFn func() ([]*habitapi.Habit, error)
    ListInvoked bool

    CreateFn func(dto habitapi.CreateHabitDTO) (*habitapi.Habit, error)
    CreateInvoked bool

    UpdateFn func(id string, dto habitapi.UpdateHabitDTO) (*habitapi.Habit, error)
    UpdateInvoked bool

    DeleteFn func(id string) error
    DeleteInvoked bool
}

func (s *HabitService) GetById(id string) (*habitapi.Habit, error) {
    s.GetByIdInvoked = true
    return s.GetByIdFn(id)
}

func (s *HabitService) List() ([]*habitapi.Habit, error) {
    s.ListInvoked = true
    return s.ListFn()
}

func (s *HabitService) Create(dto habitapi.CreateHabitDTO) (*habitapi.Habit, error) {
    s.CreateInvoked = true
    return s.CreateFn(dto)
}

func (s *HabitService) Update(id string, dto habitapi.UpdateHabitDTO) (*habitapi.Habit, error) {
    s.UpdateInvoked = true
    return s.UpdateFn(id, dto)
}

func (s *HabitService) Delete(id string) error {
    s.DeleteInvoked = true
    return s.DeleteFn(id)
}
