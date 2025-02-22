package mongodb

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/kperilla/habitapi/habitapi"
)

type RewardService struct {
	DB *mongo.Database
    CollectionName string
}

func (s *RewardService) GetById(id string) (*habitapi.Reward, error) {
    empty := &habitapi.Reward{}
    reward, err := GetById(id, s.CollectionName, empty, s.DB)
    return reward, err
}

func (s * RewardService) List() ([]*habitapi.Reward, error) {
    empty := []*habitapi.Reward{}
    rewards, err := List(s.CollectionName, empty, s.DB)
    return rewards, err
}

func (s *RewardService) Create(dto habitapi.CreateRewardDTO) (*habitapi.Reward, error) {
    rewardRaw, id, err := Create(&dto, s.CollectionName, s.DB)
    reward := rewardRaw.(habitapi.Reward)
    reward.ID = id
    return &reward, err
}

func (s *RewardService) Update(
    id string, dto habitapi.UpdateRewardDTO,
) (*habitapi.Reward, error) {
    rewardRaw, err := Update(id, &dto, s.CollectionName, s.DB)
    reward := rewardRaw.(habitapi.Reward)
    return &reward, err
}

func (s *RewardService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
