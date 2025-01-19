package mongodb

import (
	"log"
	"os"
	"testing"

	"github.com/kperilla/habitapi/habitapi"
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
    var dtoA = habitapi.CreateUserDTO{Name: "test"}
    user, err := userService.CreateUser(dtoA)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }
    createdId := user.ID

    var dtoB = habitapi.CreateUserDTO{Name: "test2"}
    user2, err := userService.CreateUser(dtoB)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }
    createdId2 := user2.ID

    // Get
    retrievedUser, err := userService.User(createdId)
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
    err = userService.DeleteUser(createdId)
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
