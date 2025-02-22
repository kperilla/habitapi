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
    collectionName := "users"

    // Create
    var dtoA = habitapi.CreateUserDTO{Name: "test"}
    user1Raw, id1, err := Create(&dtoA, collectionName, db)
    user1 := user1Raw.(habitapi.User)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }
    createdId := id1

    var dtoB = habitapi.CreateUserDTO{Name: "test2"}
    user2Raw, id2, err := Create(&dtoB, collectionName, db)
    user2 := user2Raw.(habitapi.User)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Create failed")
    }
    createdId2 := id2

    // Update
    newName := "ChangedTest"
    var updateDto = habitapi.UpdateUserDTO{Name: newName}
    changedUserRaw, err := Update(createdId2.Hex(), &updateDto, collectionName, db)
    changedUser := changedUserRaw.(habitapi.User)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Update failed")
    }
    if changedUser.Name != newName {
        t.Errorf("Expected user name %s, got %s", newName, changedUser.Name)
    }

    // Get
    emptyUser := &habitapi.User{}
    retrievedUserRaw, err := GetById(createdId.Hex(), collectionName, emptyUser, db)
    retrievedUser := retrievedUserRaw.(*habitapi.User)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get failed")
    }
    if retrievedUser.Name != user1.Name {
        t.Errorf("Expected user name %s, got %s", user1.Name, retrievedUser.Name)
    }
    emptyUser = &habitapi.User{}
    retrievedUser2Raw, err := GetById(createdId2.Hex(), collectionName, emptyUser, db)
    retrievedUser2 := retrievedUser2Raw.(*habitapi.User)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get failed")
    }
    if retrievedUser2.Name == user2.Name {
        t.Errorf("Expected user name %s, got %s", newName, retrievedUser2.Name)
    }

    // Get All
    emptyUserList := []*habitapi.User{}
    users, err := List(collectionName, emptyUserList, db)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get all failed")
    }
    if len(users) != 2 {
        t.Errorf("Expected 2 users, got %d", len(users))
    }

    // Delete
    err = Delete(createdId.Hex(), collectionName, db)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Delete failed")
    }
    err = Delete(createdId2.Hex(), collectionName, db)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Delete failed")
    }
}
