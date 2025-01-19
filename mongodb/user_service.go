package mongodb

import (
    "log"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "github.com/kperilla/habitapi/habitapi"
)

type UserService struct {
	DB *mongo.Database
}

func (s *UserService) User(id string) (*habitapi.User, error) {
    objectId, _ := bson.ObjectIDFromHex(id)
    user := &habitapi.User{}
    result := s.DB.Collection("users").FindOne(nil, bson.M{"_id": objectId})
    err := result.Err()
    if err == mongo.ErrNoDocuments {
        return nil, &habitapi.ErrUserNotFound{Err: err}
    }
    err = result.Decode(user)
    if err != nil {
        log.Fatal(err)
    }
    return user, err
}

func (s *UserService) Users() ([]*habitapi.User, error) {
    users := []*habitapi.User{}
    cursor, err := s.DB.Collection("users").Find(nil, bson.D{})
    if err != nil {
        log.Fatal(err)
    }
    err = cursor.All(nil, &users)
    if err != nil {
        log.Fatal(err)
    }
    return users, err
}

func (s *UserService) CreateUser(dto habitapi.CreateUserDTO) (*habitapi.User, error) {
    user := dto.ToModel()
    res, err := s.DB.Collection("users").InsertOne(nil, user)
    if err != nil {
        log.Fatal(err)
    }
    user.ID = res.InsertedID.(bson.ObjectID).Hex()
    return &user, err
}

func (s *UserService) DeleteUser(id string) error {
    objectId, _ := bson.ObjectIDFromHex(id)
    _, err := s.DB.Collection("users").DeleteOne(nil, bson.M{"_id": objectId})
    return err
}
