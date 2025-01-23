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
    user, err := GetById(id, s.CollectionName, empty, s.DB)
    return user, err
}

func (s * UserService) List() ([]*habitapi.User, error) {
    empty := []*habitapi.User{}
    users, err := List(s.CollectionName, empty, s.DB)
    return users, err
}

func (s *UserService) Create(dto habitapi.CreateUserDTO) (*habitapi.User, error) {
    user, id, err := Create(&dto, s.CollectionName, s.DB)
    user.ID = id
    return user, err
}

func (s *UserService) Update(
    id string, dto habitapi.UpdateUserDTO,
) (*habitapi.User, error) {
    user, err := Update(id, &dto, s.CollectionName, s.DB)
    return user, err
}

func (s *UserService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
