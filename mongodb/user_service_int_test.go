package mongodb

import (
    "log"
    "os"
    "testing"

    // "go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetCreateUserIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }

    mongo_uri := os.Getenv("MONGODB_URI")
    mongo_username := os.Getenv("MONGODB_USERNAME")
    mongo_password := os.Getenv("MONGODB_PASSWORD")
    credential := Credential{
        Username: mongo_username,
        Password: mongo_password,
    }
    client, err := InitMongo(mongo_uri, credential)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := client.Disconnect(nil); err != nil {
            panic(err)
        }
    }()
    db := client.Database("test")

    // Create
    userService := &UserService{DB: db}
    user, created_id, err := userService.CreateUser("test")
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }

    _, createdId2, err := userService.CreateUser("test2")
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }

    // Get
    retrievedUser, err := userService.User(created_id)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get failed")
    }
    if retrievedUser.Name != user.Name {
        t.Errorf("Expected user name %s, got %s", user.Name, retrievedUser.Name)
    }

    // Get All
    users, err := userService.Users()
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get all failed")
    }
    if len(users) != 2 {
        t.Errorf("Expected 2 users, got %d", len(users))
    }

    // Delete
    err = userService.DeleteUser(created_id)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Delete failed")
    }
    err = userService.DeleteUser(createdId2)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Delete failed")
    }
}
