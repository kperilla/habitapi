package mongodb

import (
    // "context"
    // "time"
    "log"

    "go.mongodb.org/mongo-driver/v2/mongo"
    // "go.mongodb.org/mongo-driver/v2/bson"
    "habits/habittrackerapi"
)

type UserService struct {
	DB *mongo.Database
}

func (s *UserService) User(name string) (*habittrackerapi.User, error) {
    user := &habittrackerapi.User{}
    // log.Println(user)
    result := s.DB.Collection("users").FindOne(nil, bsonFilter("name", name))
    // log.Println(result.Raw())
    // id := result.Raw().(bson.M)["id"].(string)
    err := result.Decode(user)
    if err != nil {
        log.Fatal(err)
    }
    // log.Println(*user)
    return user, err
}

func (s *UserService) CreateUser(name string) (*habittrackerapi.User, error) {
    // ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    // defer cancel()
    _, err := s.DB.Collection("users").InsertOne(nil, bsonFilter("name", name))
    if err != nil {
        log.Fatal(err)
    }
    // var id bson.ObjectID = res.InsertedID.(bson.ObjectID)
    user := habittrackerapi.User{Name: name}
    return &user, err
}
