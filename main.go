package main

import (
    "log"
    "context"
    "time"
    "os"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
    "habits/http"
)

func init_mongo(mongo_uri string) (*mongo.Client, context.Context, context.CancelFunc) {
    client, _ := mongo.Connect(options.Client().ApplyURI(mongo_uri))
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    err := client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal(err)
    }
    return client, ctx, cancel
}

func main() {
    mongo_uri := os.Getenv("MONGO_URI")
    client, ctx, cancel := init_mongo(mongo_uri)
    defer cancel()

    server := http.NewAPIServer(":8080")
    var handler http.Handler
    if err := server.Run(&handler); err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()
}
