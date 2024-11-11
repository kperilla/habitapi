package main

import (
    "log"
    "os"

    "github.com/kperilla/habitapi/http"
    "github.com/kperilla/habitapi/mongodb"
)

func main() {
    mongo_uri := os.Getenv("MONGO_URI")
    mongo_username := os.Getenv("MONGODB_USERNAME")
    mongo_password := os.Getenv("MONGODB_PASSWORD")
    credential := mongodb.Credential{
        Username: mongo_username,
        Password: mongo_password,
    }
    client, err := mongodb.InitMongo(mongo_uri, credential)
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("habits")
    _ = &mongodb.UserService{DB: db}

    server := http.NewAPIServer(":8080")
    userService := &mongodb.UserService{DB: db}
    // handler := http.Handler{UserService: userService}
    handler := http.NewHandler(userService)
    if err := server.Run(handler); err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := client.Disconnect(nil); err != nil {
            panic(err)
        }
    }()
}
