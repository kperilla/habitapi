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

    mongo_uri := os.Getenv("MONGO_URI")
    mongo_password := os.Getenv("MONGO_LOCAL_PASSWORD")
    credential := Credential{
        Username: "mongo-admin",
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

    // Get
    retrievedUser, err := userService.User(created_id)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Get failed")
    }
    if retrievedUser.Name != user.Name {
        t.Errorf("Expected user name %s, got %s", user.Name, retrievedUser.Name)
    }

    // Delete
    err = userService.DeleteUser(created_id)
    if err != nil {
        log.Fatal(err)
        t.Errorf("Delete failed")
    }
}

// func TestUser(t *testing.T) {
//     db := mock.Database{}
//     collection := mock.Collection{}
//     db.CollectionFn = func(name string) *mongodb.Collection {
//         return &collection
//     }
//     collection.FindOneFn = func(ctx mock.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
//         if filter != mongodb.bsonFilter("id", "1") {
//             t.Fatalf("unexpected id: %d", filter)
//         }
//         return &mongo.SingleResult{}
//     }
//     userService := &mongodb.UserService{DB: db}
//
//     user, _ := userService.User("1")
//     if user.Id != "1" {
//         t.Errorf("Expected user id 1, got %s", user.Id)
//     }
// }
//
// func TestCreateUser(t *testing.T) {
//     t.Skip("Skipping test")
// }


// Something tells me that the complexity of the mocks and testing is a hint
// that I'm doing something wrong. This should probably be the realm of
// integration testing.
