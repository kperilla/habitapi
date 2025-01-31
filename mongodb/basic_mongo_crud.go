package mongodb

import (
	"log"

	"github.com/kperilla/habitapi/habitapi"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetById[T any](
    id string,
    collectionName string,
    empty *T,
    db *mongo.Database,
) (*T, error) {
    objectId, _ := bson.ObjectIDFromHex(id)
    resource := empty
    result := db.Collection(collectionName).FindOne(nil, bson.M{"_id": objectId})
    err := result.Err()
    if err == mongo.ErrNoDocuments {
        return nil, &habitapi.ErrResourceNotFound{Err: err}
    }
    err = result.Decode(resource)
    if err != nil {
        log.Fatal(err)
    }
    return resource, err
}

func List[T any](
    collectionName string,
    empty []*T,
    db *mongo.Database,
) ([]*T, error) {
    resources := empty
    cursor, err := db.Collection(collectionName).Find(nil, bson.D{})
    if err != nil {
        log.Fatal(err)
    }
    err = cursor.All(nil, &resources)
    if err != nil {
        log.Fatal(err)
    }
    return resources, err
}

func Create[T any, DTO habitapi.DTO[T]] (
    dto DTO,
    collectionName string,
    db *mongo.Database,
) (*T, bson.ObjectID, error) {
    resource := dto.ToModel()
    res, err := db.Collection(collectionName).InsertOne(nil, resource)
    if err != nil {
        log.Fatal(err)
    }
    id := res.InsertedID.(bson.ObjectID)
    return &resource, id, err
}

func Update[T any, DTO habitapi.DTO[T]](
    id string,
    dto DTO,
    collectionName string,
    db *mongo.Database,
) (*T, error) {
    objectId, _ := bson.ObjectIDFromHex(id)
    filter := bson.D{{"_id", objectId}}
    update := bson.D{{"$set", dto}}
    _, err := db.Collection(collectionName).UpdateOne(nil, filter, update)
    if err != nil {
        log.Fatal(err)
    }
    // TODO: THIS IS WRONG!
    resource := dto.ToModel()


    return &resource, err
}

func Delete (
    id string,
    collectionName string,
    db *mongo.Database,
) error {
    objectId, _ := bson.ObjectIDFromHex(id)
    _, err := db.Collection(collectionName).DeleteOne(nil, bson.M{"_id": objectId})
    return err
}
