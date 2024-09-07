package mongodb

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Credential struct {
    Username string
    Password string
}

func InitMongo(mongo_uri string, credential Credential) (*mongo.Client, error) {
    mongo_creds := options.Credential{
        Username: credential.Username,
        Password: credential.Password,
    }
    client, err := mongo.Connect(options.Client().ApplyURI(mongo_uri).SetAuth(mongo_creds))
    if err != nil {
        return nil, err
    }
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())
    return client, err
}

// Credit: https://www.mongodb.com/community/forums/t/mongodb-go-primative-e/168870
func bsonFilter(key string, value string) bson.D {
	return bson.D{{Key: key, Value: value}}
}

