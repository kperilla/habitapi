package mongodb

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/kperilla/habitapi/habitapi"
)

type HabitGroupService struct {
	DB *mongo.Database
    CollectionName string
}

func (s *HabitGroupService) GetById(id string) (*habitapi.HabitGroup, error) {
    empty := &habitapi.HabitGroup{}
    group, err := GetById(id, s.CollectionName, empty, s.DB)
    return group, err
}

func (s * HabitGroupService) List() ([]*habitapi.HabitGroup, error) {
    empty := []*habitapi.HabitGroup{}
    groups, err := List(s.CollectionName, empty, s.DB)
    return groups, err
}

func (s *HabitGroupService) Create(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error) {
    groupRaw, id, err := Create(&dto, s.CollectionName, s.DB)
    group := groupRaw.(habitapi.HabitGroup)
    group.ID = id
    return &group, err
}

func (s *HabitGroupService) Update(
    id string, dto habitapi.UpdateHabitGroupDTO,
) (*habitapi.HabitGroup, error) {
    groupRaw, err := Update(id, &dto, s.CollectionName, s.DB)
    group := groupRaw.(habitapi.HabitGroup)
    return &group, err
}

func (s *HabitGroupService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
