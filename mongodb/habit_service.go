package mongodb

import (
	"github.com/kperilla/habitapi/habitapi"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type HabitService struct {
	DB *mongo.Database
    CollectionName string
}

func (s *HabitService) GetById(id string) (*habitapi.Habit, error) {
    empty := &habitapi.Habit{}
    habit, err := GetById(id, s.CollectionName, empty, s.DB)
    return habit, err
}

func (s * HabitService) List() ([]*habitapi.Habit, error) {
    empty := []*habitapi.Habit{}
    habits, err := List(s.CollectionName, empty, s.DB)
    return habits, err
}

func (s *HabitService) Create(dto habitapi.CreateHabitDTO) (*habitapi.Habit, error) {
    habitRaw, id, err := Create(&dto, s.CollectionName, s.DB)
    habit := habitRaw.(habitapi.Habit)
    habit.ID = id
    return &habit, err
}

func (s *HabitService) Update(
    id string, dto habitapi.UpdateHabitDTO,
) (*habitapi.Habit, error) {
    habitRaw, err := Update(id, &dto, s.CollectionName, s.DB)
    habit := habitRaw.(habitapi.Habit)
    return &habit, err
}

func (s *HabitService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
