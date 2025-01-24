package mock

import (
    "github.com/kperilla/habitapi/habitapi"
)

type RewardService struct {
    GetByIdFn func(id string) (*habitapi.Reward, error)
    GetByIdInvoked bool

    ListFn func() ([]*habitapi.Reward, error)
    ListInvoked bool

    CreateFn func(dto habitapi.CreateRewardDTO) (*habitapi.Reward, error)
    CreateInvoked bool

    UpdateFn func(id string, dto habitapi.UpdateRewardDTO) (*habitapi.Reward, error)
    UpdateInvoked bool

    DeleteFn func(id string) error
    DeleteInvoked bool
}

func (s *RewardService) GetById(id string) (*habitapi.Reward, error) {
    s.GetByIdInvoked = true
    return s.GetByIdFn(id)
}

func (s *RewardService) List() ([]*habitapi.Reward, error) {
    s.ListInvoked = true
    return s.ListFn()
}

func (s *RewardService) Create(dto habitapi.CreateRewardDTO) (*habitapi.Reward, error) {
    s.CreateInvoked = true
    return s.CreateFn(dto)
}

func (s *RewardService) Update(id string, dto habitapi.UpdateRewardDTO) (*habitapi.Reward, error) {
    s.UpdateInvoked = true
    return s.UpdateFn(id, dto)
}

func (s *RewardService) Delete(id string) error {
    s.DeleteInvoked = true
    return s.DeleteFn(id)
}
