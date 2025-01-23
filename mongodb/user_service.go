package mongodb

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/kperilla/habitapi/habitapi"
)

type UserService struct {
	DB *mongo.Database
    CollectionName string
}

func (s *UserService) GetById(id string) (*habitapi.User, error) {
    empty := &habitapi.User{}
    group, err := GetById(id, s.CollectionName, empty, s.DB)
    return group, err
}

func (s * UserService) List() ([]*habitapi.User, error) {
    empty := []*habitapi.User{}
    groups, err := List(s.CollectionName, empty, s.DB)
    return groups, err
}

func (s *UserService) Create(dto habitapi.CreateUserDTO) (*habitapi.User, error) {
    group, id, err := Create(&dto, s.CollectionName, s.DB)
    group.ID = id
    return group, err
}

func (s *UserService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
