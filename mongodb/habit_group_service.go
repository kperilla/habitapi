package mongodb

import (
    "log"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "github.com/kperilla/habitapi/habitapi"
)

type HabitGroupService struct {
	DB *mongo.Database
}

func (s *HabitGroupService) GetById(id string) (*habitapi.HabitGroup, error) {
    objectId, _ := bson.ObjectIDFromHex(id)
    group := &habitapi.HabitGroup{}
    result := s.DB.Collection("habit_groups").FindOne(nil, bson.M{"_id": objectId})
    err := result.Err()
    if err == mongo.ErrNoDocuments {
        return nil, &habitapi.ErrResourceNotFound{Err: err}
    }
    err = result.Decode(group)
    if err != nil {
        log.Fatal(err)
    }
    return group, err
}

func (s * HabitGroupService) List() ([]*habitapi.HabitGroup, error) {
    groups := []*habitapi.HabitGroup{}
    cursor, err := s.DB.Collection("habit_groups").Find(nil, bson.D{})
    if err != nil {
        log.Fatal(err)
    }
    err = cursor.All(nil, &groups)
    if err != nil {
        log.Fatal(err)
    }
    return groups, err
}

func (s *HabitGroupService) Create(dto habitapi.CreateHabitGroupDTO) (*habitapi.HabitGroup, error) {
    group := dto.ToModel()
    res, err := s.DB.Collection("habit_groups").InsertOne(nil, group)
    if err != nil {
        log.Fatal(err)
    }
    group.ID = res.InsertedID.(bson.ObjectID).Hex()
    return &group, err
}

func (s *HabitGroupService) Delete(id string) error {
    objectId, _ := bson.ObjectIDFromHex(id)
    _, err := s.DB.Collection("habit_groups").DeleteOne(nil, bson.M{"_id": objectId})
    return err
}
