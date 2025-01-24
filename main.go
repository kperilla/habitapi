package main

import (
    "log"
    "os"

    "github.com/kperilla/habitapi/http"
    "github.com/kperilla/habitapi/mongodb"
)

func main() {
    mongo_uri := os.Getenv("MONGODB_URI")
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
    // _ = &mongodb.UserService{DB: db}

    server := http.NewAPIServer(":8080")
    userService := &mongodb.UserService{DB: db, CollectionName: "users"}
    habitGroupService := &mongodb.HabitGroupService{DB: db, CollectionName: "habit_groups"}
    habitService := &mongodb.HabitService{DB: db, CollectionName: "habits"}
    deedService := &mongodb.DeedService{DB: db, CollectionName: "deeds"}
    rewardService := &mongodb.RewardService{DB: db, CollectionName: "rewards"}
    routeHandler := http.NewHandler(
        userService, habitGroupService, habitService, deedService, rewardService,
    )
    err = server.Run(
        routeHandler,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := client.Disconnect(nil); err != nil {
            panic(err)
        }
    }()
}
