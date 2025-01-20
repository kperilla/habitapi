package mongodb

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/kperilla/habitapi/habitapi"
)

type HabitGroupService struct {
	DB *mongo.Database
}

func (s *HabitGroupService) GetById(id string) (*habitapi.HabitGroup, error) {
    empty := &habitapi.HabitGroup{}
    group, err := GetById(id, "habit_groups", empty, s.DB)
    return group, err
}

func (s * HabitGroupService) List() ([]*habitapi.HabitGroup, error) {
    empty := []*habitapi.HabitGroup{}
    groups, err := List("habit_groups", empty, s.DB)
    return groups, err
}

func (s *HabitGroupService) Create(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error) {
    group, id, err := Create(&dto, "habit_groups", s.DB)
    group.ID = id
    return group, err
}

func (s *HabitGroupService) Delete(id string) error {
    err := Delete(id, "habit_groups", s.DB)
    return err
}
