package mongodb

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/kperilla/habitapi/habitapi"
)

type DeedService struct {
	DB *mongo.Database
    CollectionName string
}

func (s *DeedService) GetById(id string) (*habitapi.Deed, error) {
    empty := &habitapi.Deed{}
    deedRaw, err := GetById(id, s.CollectionName, empty, s.DB)
    deed := deedRaw.(*habitapi.Deed)
    return deed, err
}

func (s * DeedService) List() ([]*habitapi.Deed, error) {
    empty := []*habitapi.Deed{}
    deeds, err := List(s.CollectionName, empty, s.DB)
    return deeds, err
}

func (s *DeedService) Create(dto habitapi.CreateDeedDTO) (*habitapi.Deed, error) {
    deedRaw, id, err := Create(&dto, s.CollectionName, s.DB)
    deed := deedRaw.(habitapi.Deed)
    deed.ID = id
    return &deed, err
}

func (s *DeedService) Update(
    id string, dto habitapi.UpdateDeedDTO,
) (*habitapi.Deed, error) {
    deedRaw, err := Update(id, &dto, s.CollectionName, s.DB)
    deed := deedRaw.(habitapi.Deed)
    return &deed, err
}

func (s *DeedService) Delete(id string) error {
    err := Delete(id, s.CollectionName, s.DB)
    return err
}
