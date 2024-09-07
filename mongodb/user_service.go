package mongodb

import (
    // "context"
    // "time"
    "log"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "habits/habittrackerapi"
)

type UserService struct {
	DB *mongo.Database
}

func (s *UserService) User(id string) (*habittrackerapi.User, error) {
    objectId, _ := bson.ObjectIDFromHex(id)
    user := &habittrackerapi.User{}
    result := s.DB.Collection("users").FindOne(nil, bson.M{"_id": objectId})
    err := result.Decode(user)
    if err != nil {
        log.Fatal(err)
    }
    return user, err
}

func (s *UserService) CreateUser(name string) (*habittrackerapi.User, string, error) {
    // ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    // defer cancel()
    res, err := s.DB.Collection("users").InsertOne(nil, bsonFilter("name", name))
    if err != nil {
        log.Fatal(err)
    }
    id :=res.InsertedID.(bson.ObjectID).Hex()
    user := habittrackerapi.User{Name: name}
    return &user, id, err
}

func (s *UserService) DeleteUser(id string) error {
    objectId, _ := bson.ObjectIDFromHex(id)
    _, err := s.DB.Collection("users").DeleteOne(nil, bson.M{"_id": objectId})
    return err
}
